package cover

import (
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

func drawDistributedForms(dc *gg.Context, formType int) {
	margin := 80.0
	gridSize := 5
	cellW := (float64(width) - margin*2) / float64(gridSize)
	cellH := (float64(height) - margin*2) / float64(gridSize*3)

	numForms := 15 + rand.Intn(10)

	dc.SetColor(oklchToRGB(formColor))

	for i := 0; i < numForms; i++ {
		gridX := rand.Intn(gridSize)
		gridY := rand.Intn(gridSize * 3)

		baseX := margin + float64(gridX)*cellW + cellW/2
		baseY := margin + float64(gridY)*cellH + cellH/2

		jitterX := (rand.Float64() - 0.5) * cellW * 0.8
		jitterY := (rand.Float64() - 0.5) * cellH * 0.8

		x := baseX + jitterX
		y := baseY + jitterY

		size := 150 + rand.Float64()*200
		radius := size / 2

		drawFormAt(dc, formType, x, y, radius)
	}
}

func drawFormAt(dc *gg.Context, formType int, x, y, size float64) {
	margin := 60.0
	maxW := float64(width) - margin
	maxH := float64(height) - margin

	switch formType {
	case 0:
		if x-size < margin || x+size > maxW || y-size < margin || y+size > maxH {
			return
		}
		dc.DrawCircle(x, y, size)
		dc.Fill()
	case 1:
		if x-size < margin || x+size > maxW || y-size < margin || y+size > maxH {
			return
		}
		dc.Push()
		dc.RotateAbout(rand.Float64()*2*math.Pi, x, y)
		dc.DrawRectangle(x-size, y-size, size*2, size*2)
		dc.Fill()
		dc.Pop()
	case 2:
		h := size * math.Sqrt(3)
		if x-h < margin || x+h > maxW || y-h < margin || y+h > maxH {
			return
		}
		dc.Push()
		dc.RotateAbout(rand.Float64()*2*math.Pi, x, y)
		dc.MoveTo(x, y-h)
		dc.LineTo(x-h, y+h)
		dc.LineTo(x+h, y+h)
		dc.ClosePath()
		dc.Fill()
		dc.Pop()
	case 3:
		dc.SetLineWidth(8 + rand.Float64()*20)
		x2 := x + (rand.Float64()-0.5)*size*3
		y2 := y + (rand.Float64()-0.5)*size*3
		if x < margin || x > maxW || y < margin || y > maxH ||
			x2 < margin || x2 > maxW || y2 < margin || y2 > maxH {
			return
		}
		dc.DrawLine(x, y, x2, y2)
		dc.Stroke()
	case 4:
		numPuntos := 2 + rand.Intn(5)
		for i := 0; i < numPuntos; i++ {
			px := x + (rand.Float64()-0.5)*size*2
			py := y + (rand.Float64()-0.5)*size*2
			pr := 15 + rand.Float64()*35
			if px-pr < margin || px+pr > maxW || py-pr < margin || py+pr > maxH {
				continue
			}
			dc.DrawCircle(px, py, pr)
			dc.Fill()
		}
	}
}

func drawRandomForm(dc *gg.Context, formType int) {
	margin := 80.0
	minPos := margin
	maxPosW := float64(width) - margin
	maxPosH := float64(height) - margin

	x := minPos + rand.Float64()*(maxPosW-minPos)
	y := minPos + rand.Float64()*(maxPosH-minPos)

	size := 150 + rand.Float64()*200

	drawFormAt(dc, formType, x, y, size/2)
}
