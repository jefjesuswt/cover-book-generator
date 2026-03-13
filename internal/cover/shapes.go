package cover

import (
	"math"
	"math/rand"

	"github.com/tdewolff/canvas"
)

func drawDistributedForms(ctx *canvas.Context, formType int) {
	margin := 80.0
	gridCols := 5
	gridRows := 15
	cellW := (width - margin*2) / float64(gridCols)
	cellH := (height - margin*2) / float64(gridRows)

	type cell struct {
		x, y int
	}
	var cells []cell

	for r := range gridRows {
		for c := range gridCols {
			cells = append(cells, cell{c, r})
		}
	}

	// Mezclar las celdas para una distribución uniforme pero aleatoria
	rand.Shuffle(len(cells), func(i, j int) {
		cells[i], cells[j] = cells[j], cells[i]
	})

	numForms := 15 + rand.Intn(10)
	if numForms > len(cells) {
		numForms = len(cells)
	}

	ctx.SetFillColor(oklchToRGB(formColor))

	for i := 0; i < numForms; i++ {
		c := cells[i]
		baseX := margin + float64(c.x)*cellW + cellW/2
		baseY := margin + float64(c.y)*cellH + cellH/2

		jitterX := (rand.Float64() - 0.5) * cellW * 0.7
		jitterY := (rand.Float64() - 0.5) * cellH * 0.7

		x := baseX + jitterX
		y := baseY + jitterY

		size := 100 + rand.Float64()*150
		radius := size / 2

		drawFormAt(ctx, formType, x, y, radius)
	}
}

func drawFormAt(ctx *canvas.Context, formType int, x, y, size float64) {
	margin := 60.0
	maxW := width - margin
	maxH := height - margin

	// Ajustar Y para canvas (Y-up)
	canvasY := height - y

	switch formType {
	case 0: // Círculo
		if x-size < margin || x+size > maxW || y-size < margin || y+size > maxH {
			return
		}
		ctx.DrawPath(x, canvasY, canvas.Circle(size))
	case 1: // Cuadrado rotado
		if x-size < margin || x+size > maxW || y-size < margin || y+size > maxH {
			return
		}
		p := canvas.Rectangle(size*2, size*2)
		p = p.Transform(canvas.Identity.Rotate(rand.Float64() * 360).Translate(-size, -size))
		ctx.DrawPath(x, canvasY, p)
	case 2: // Triángulo
		h := size * math.Sqrt(3)
		if x-h < margin || x+h > maxW || y-h < margin || y+h > maxH {
			return
		}
		p := &canvas.Path{}
		p.MoveTo(0, h)
		p.LineTo(-h, -h)
		p.LineTo(h, -h)
		p.Close()
		p = p.Transform(canvas.Identity.Rotate(rand.Float64() * 360))
		ctx.DrawPath(x, canvasY, p)
	case 3: // Línea
		ctx.SetStrokeWidth(8 + rand.Float64()*10) // Reduciendo un poco el ancho comparado con gg
		ctx.SetStrokeColor(oklchToRGB(formColor))
		
		x2 := x + (rand.Float64()-0.5)*size*3
		y2 := y + (rand.Float64()-0.5)*size*3
		
		if x < margin || x > maxW || y < margin || y > maxH ||
			x2 < margin || x2 > maxW || y2 < margin || y2 > maxH {
			return
		}
		
		p := &canvas.Path{}
		p.MoveTo(x, canvasY)
		p.LineTo(x2, height-y2)
		ctx.DrawPath(0, 0, p)
	case 4: // Multi-puntos
		numPuntos := 2 + rand.Intn(5)
		for i := 0; i < numPuntos; i++ {
			px := x + (rand.Float64()-0.5)*size*2
			py := y + (rand.Float64()-0.5)*size*2
			pr := 15 + rand.Float64()*35
			if px-pr < margin || px+pr > maxW || py-pr < margin || py+pr > maxH {
				continue
			}
			ctx.DrawPath(px, height-py, canvas.Circle(pr))
		}
	}
}
