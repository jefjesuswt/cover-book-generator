package cover

import (
	"fmt"
	"image/color"
	"os"

	"github.com/tdewolff/canvas"
)

func drawText(ctx *canvas.Context, title, author string, fontPath string, textCol, authorCol color.RGBA) error {
	if _, err := os.Stat(fontPath); err != nil {
		return fmt.Errorf("fuente no encontrada: %s", fontPath)
	}

	fontFamily := canvas.NewFontFamily("theme")
	if err := fontFamily.LoadFontFile(fontPath, canvas.FontRegular); err != nil {
		return err
	}

	margin := 120.0
	maxWidth := width - margin*2

	fontSize := 600.0

	renderTitle := func(fs float64) *canvas.Text {
		face := fontFamily.Face(fs, textCol, canvas.FontRegular, canvas.FontNormal)
		return canvas.NewTextBox(face, title, maxWidth, 0.0, canvas.Center, canvas.Top, nil)
	}

	var textTitle *canvas.Text
	for {
		textTitle = renderTitle(fontSize)
		bounds := textTitle.Bounds()
		if (bounds.W() <= maxWidth+5 && bounds.H() <= height*0.45) || fontSize < 60 {
			break
		}
		fontSize -= 10
	}

	titleY := height * 0.82
	ctx.DrawText(margin, titleY, textTitle)

	authorFontSize := fontSize * 0.55
	if authorFontSize < 100 {
		authorFontSize = 100
	}

	authorFace := fontFamily.Face(authorFontSize, authorCol, canvas.FontRegular, canvas.FontNormal)
	textAuthor := canvas.NewTextBox(authorFace, author, maxWidth, 0.0, canvas.Center, canvas.Top, nil)

	authorY := height * 0.15
	ctx.DrawText(margin, authorY, textAuthor)

	return nil
}
