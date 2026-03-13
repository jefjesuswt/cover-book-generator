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

func Generate(title, author, path string) error {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crear lienzo y contexto
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	// En canvas (Y hacia arriba), invertimos la vista para trabajar con coordenadas top-down si preferimos
	// o simplemente adaptamos las funciones. Vamos a usar la convención de canvas de Y hacia arriba.
	
	// 1. Generar fondo (solo color sólido)
	backgroundColor := randomOKLCH()
	col := oklchToRGB(backgroundColor)
	
	ctx.SetFillColor(col)
	ctx.DrawPath(0, 0, canvas.Rectangle(width, height))

	// 2. Generar colores para formas y texto
	bgColors := []Oklch{backgroundColor, backgroundColor}
	initColors(bgColors, 0)

	// 3. Dibujar formas aleatorias con distribución mejorada
	globalFormType := rand.Intn(5)
	drawDistributedForms(ctx, globalFormType)

	// 4. Dibujar texto
	err := drawText(ctx, title, author, backgroundColor)
	if err != nil {
		return err
	}

	// Guardar a PNG (3.78 DPMM es aprox 96 DPI, para 1600x2560 mantenemos la escala 1:1)
	return renderers.Write(path, c, canvas.DPMM(1.0))
}
