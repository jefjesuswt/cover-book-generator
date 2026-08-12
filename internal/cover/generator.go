package cover

import (
	"math/rand"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

const (
	width  = 1600.0
	height = 2560.0
)

func Generate(title, author, path, apiKey string) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	theme := detectTheme(title)

	// Try image background
	img, err := fetchBackground(apiKey, theme.Query)
	if err == nil {
		ctx.DrawImage(0, 0, img, canvas.DPMM(1.0))
		colors := extractColors(img)

		overlayAlpha := uint8(theme.Overlay * 255)
		overlayCol := colors.overlay
		overlayCol.A = overlayAlpha
		ctx.SetFillColor(overlayCol)
		ctx.DrawPath(0, 0, canvas.Rectangle(width, height))

		if err := drawText(ctx, title, author, theme.FontFile, colors.text, colors.textMuted); err != nil {
			return err
		}
	} else {
		// Fallback: random OKLCH background
		bgColor := randomOKLCH()
		col := oklchToRGB(bgColor)
		ctx.SetFillColor(col)
		ctx.DrawPath(0, 0, canvas.Rectangle(width, height))

		initColors([]Oklch{bgColor, bgColor}, 0)
		fallbackText := oklchToRGB(titleColor)
		fallbackAuthor := oklchToRGB(authorColor)
		if err := drawText(ctx, title, author, theme.FontFile, fallbackText, fallbackAuthor); err != nil {
			return err
		}
	}

	return renderers.Write(path, c, canvas.DPMM(1.0))
}
