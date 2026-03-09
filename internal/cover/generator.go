package cover

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	width  = 1600
	height = 2560
)

func Generate(title, author, path string) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crear contexto
	dc := gg.NewContext(width, height)

	// 1. Generar fondo (solo color sólido)
	backgroundColor := randomOKLCH()
	col := oklchToRGB(backgroundColor)
	dc.SetColor(col)
	dc.Clear()

	// 2. Generar colores para formas y texto
	bgColors := []Oklch{backgroundColor, backgroundColor}
	initColors(bgColors, 0)

	// 3. Dibujar formas aleatorias con distribución mejorada
	globalFormType := rand.Intn(5)
	drawDistributedForms(dc, globalFormType)

	// 4. Dibujar texto
	err := drawText(dc, title, author, backgroundColor)
	if err != nil {
		return err
	}

	// Guardar
	return dc.SavePNG(path)
}
