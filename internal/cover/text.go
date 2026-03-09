package cover

import (
	"errors"
	"os"

	"github.com/fogleman/gg"
)

func drawText(dc *gg.Context, title, author string, bgColor Oklch) error {
	fontPath := "fonts/Roboto-Regular.ttf"

	if _, err := os.Stat(fontPath); err != nil {
		return errors.New("fuente no encontrada: " + fontPath)
	}

	margin := 80.0
	maxWidth := float64(width) - margin*2

	fontSize := 280.0

	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return err
	}

	dc.SetColor(oklchToRGB(titleColor))

	for {
		_, h := dc.MeasureMultilineString(title, 1.15)
		if h <= float64(height)*0.4 || fontSize < 40 {
			break
		}
		fontSize -= 10
		if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
			return err
		}
	}

	titleAreaHeight := float64(height) * 0.45
	titleY := (float64(height) - titleAreaHeight) / 2

	dc.DrawStringWrapped(title, float64(width)/2, titleY, 0.5, 0, maxWidth, 1.15, gg.AlignCenter)

	dc.SetColor(oklchToRGB(authorColor))
	authorFontSize := fontSize * 0.35
	if authorFontSize < 35 {
		authorFontSize = 35
	}
	if authorFontSize > 80 {
		authorFontSize = 80
	}

	if err := dc.LoadFontFace(fontPath, authorFontSize); err != nil {
		return err
	}

	authorMarginBottom := 100.0
	authorY := float64(height) - authorMarginBottom
	dc.DrawStringWrapped(author, float64(width)/2, authorY, 0.5, 1, maxWidth, 1.0, gg.AlignCenter)

	return nil
}
