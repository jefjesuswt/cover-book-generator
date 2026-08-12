package cover

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"math"
	"net/http"
	"net/url"

	"golang.org/x/image/draw"
)

type unsplashResponse struct {
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

// fetchImageFromUnsplash searches Unsplash for a random image matching the query.
func fetchImageFromUnsplash(apiKey, query string) (image.Image, error) {
	endpoint := fmt.Sprintf(
		"https://api.unsplash.com/photos/random?query=%s&orientation=portrait",
		url.QueryEscape(query),
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Client-ID "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unsplash API returned %d", resp.StatusCode)
	}

	var ur unsplashResponse
	if err := json.NewDecoder(resp.Body).Decode(&ur); err != nil {
		return nil, err
	}
	if ur.URLs.Regular == "" {
		return nil, fmt.Errorf("no image URL in response")
	}

	imgResp, err := http.Get(ur.URLs.Regular)
	if err != nil {
		return nil, err
	}
	defer imgResp.Body.Close()

	img, err := jpeg.Decode(imgResp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// resizeAndCrop scales image to fill target dimensions, center-cropping excess.
func resizeAndCrop(src image.Image, targetW, targetH int) image.Image {
	srcW := src.Bounds().Dx()
	srcH := src.Bounds().Dy()

	scale := math.Max(float64(targetW)/float64(srcW), float64(targetH)/float64(srcH))
	scaledW := int(math.Round(float64(srcW) * scale))
	scaledH := int(math.Round(float64(srcH) * scale))

	scaled := image.NewRGBA(image.Rect(0, 0, scaledW, scaledH))
	draw.CatmullRom.Scale(scaled, scaled.Bounds(), src, src.Bounds(), draw.Over, nil)

	cropX := (scaledW - targetW) / 2
	cropY := (scaledH - targetH) / 2

	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	for y := 0; y < targetH; y++ {
		for x := 0; x < targetW; x++ {
			dst.Set(x, y, scaled.At(cropX+x, cropY+y))
		}
	}
	return dst
}

// fetchBackground tries theme query, then generic, returns error if both fail.
func fetchBackground(apiKey, query string) (image.Image, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("no API key")
	}

	img, err := fetchImageFromUnsplash(apiKey, query)
	if err == nil {
		return resizeAndCrop(img, int(width), int(height)), nil
	}

	img, err2 := fetchImageFromUnsplash(apiKey, "book abstract")
	if err2 == nil {
		return resizeAndCrop(img, int(width), int(height)), nil
	}

	return nil, fmt.Errorf("unsplash failed: %v, retry: %v", err, err2)
}
