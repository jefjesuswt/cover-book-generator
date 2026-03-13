package cover

import (
	"fmt"
	"os"

	"github.com/tdewolff/canvas"
)

func drawText(ctx *canvas.Context, title, author string, bgColor Oklch) error {
	fontPath := "fonts/Roboto-Regular.ttf"

	if _, err := os.Stat(fontPath); err != nil {
		return fmt.Errorf("fuente no encontrada: %s", fontPath)
	}

	fontFamily := canvas.NewFontFamily("Roboto")
	if err := fontFamily.LoadFontFile(fontPath, canvas.FontRegular); err != nil {
		return err
	}

	margin := 120.0
	maxWidth := width - margin*2

	// Algoritmo de ajuste inteligente: empezamos con un tamaño máximo razonable
	fontSize := 600.0 // 800 era demasiado

	renderTitle := func(fs float64) *canvas.Text {
		face := fontFamily.Face(fs, oklchToRGB(titleColor), canvas.FontRegular, canvas.FontNormal)
		// NewTextBox ajusta el texto al ancho maxWidth
		return canvas.NewTextBox(face, title, maxWidth, 0.0, canvas.Center, canvas.Top, nil)
	}

	var textTitle *canvas.Text
	for {
		textTitle = renderTitle(fontSize)
		bounds := textTitle.Bounds()

		// Comprobamos:
		// 1. Que no haya desbordamiento horizontal (un solo glifo o palabra muy larga)
		// 2. Que la altura total no exceda el 45% del lienzo
		if (bounds.W() <= maxWidth+5 && bounds.H() <= height*0.45) || fontSize < 60 {
			break
		}
		fontSize -= 10
	}

	// Posicionar título
	// En canvas (Y-up) con valign=Top, DrawText usa el Y como la parte SUPERIOR de la caja.
	// Queremos que el bloque esté centrado alrededor de un eje vertical a ~70% de la altura
	titleY := height * 0.82
	ctx.DrawText(margin, titleY, textTitle)

	// Renderizar el Autor con mayor presencia (aprox 55% del título)
	authorFontSize := fontSize * 0.55
	if authorFontSize < 100 {
		authorFontSize = 100
	}

	authorFace := fontFamily.Face(authorFontSize, oklchToRGB(authorColor), canvas.FontRegular, canvas.FontNormal)
	// Para el autor también usamos TextBox por si es muy largo, pero suele ser una línea
	textAuthor := canvas.NewTextBox(authorFace, author, maxWidth, 0.0, canvas.Center, canvas.Top, nil)

	// El autor se dibuja cerca del fondo (un poco más arriba para dar aire)
	authorY := height * 0.15
	ctx.DrawText(margin, authorY, textAuthor)

	return nil
}
