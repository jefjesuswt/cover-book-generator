package cover

import (
	"image/color"
	"math"
	"math/rand"
)

type Oklch struct {
	L, C, H float64
}

var (
	formColor   Oklch
	titleColor  Oklch
	authorColor Oklch
	bgColors    []Oklch
)

type TailwindColor struct {
	name  string
	light Oklch // 50, 100, 200
	dark  Oklch // 800, 900, 950
}

var tailwindPalette = []TailwindColor{
	// Warm
	{"red", Oklch{0.97, 0.06, 25}, Oklch{0.13, 0.12, 20}},
	{"orange", Oklch{0.96, 0.07, 40}, Oklch{0.14, 0.10, 30}},
	{"amber", Oklch{0.97, 0.06, 60}, Oklch{0.14, 0.09, 45}},
	{"yellow", Oklch{0.97, 0.06, 90}, Oklch{0.15, 0.08, 70}},
	// Cool
	{"lime", Oklch{0.97, 0.06, 140}, Oklch{0.14, 0.09, 120}},
	{"green", Oklch{0.96, 0.06, 160}, Oklch{0.13, 0.10, 150}},
	{"emerald", Oklch{0.95, 0.05, 170}, Oklch{0.13, 0.10, 165}},
	{"teal", Oklch{0.96, 0.05, 180}, Oklch{0.13, 0.09, 175}},
	{"cyan", Oklch{0.96, 0.05, 200}, Oklch{0.14, 0.09, 190}},
	{"sky", Oklch{0.96, 0.05, 210}, Oklch{0.13, 0.08, 205}},
	{"blue", Oklch{0.96, 0.05, 240}, Oklch{0.13, 0.10, 230}},
	{"indigo", Oklch{0.95, 0.06, 265}, Oklch{0.13, 0.10, 255}},
	{"violet", Oklch{0.95, 0.06, 290}, Oklch{0.13, 0.10, 280}},
	{"purple", Oklch{0.95, 0.06, 310}, Oklch{0.13, 0.10, 295}},
	{"fuchsia", Oklch{0.96, 0.07, 325}, Oklch{0.13, 0.11, 310}},
	{"pink", Oklch{0.96, 0.07, 340}, Oklch{0.13, 0.10, 325}},
	{"rose", Oklch{0.96, 0.07, 355}, Oklch{0.13, 0.10, 345}},
	// Neutral
	{"slate", Oklch{0.97, 0.02, 240}, Oklch{0.13, 0.03, 250}},
	{"gray", Oklch{0.97, 0.02, 260}, Oklch{0.13, 0.03, 265}},
	{"zinc", Oklch{0.97, 0.02, 280}, Oklch{0.13, 0.02, 285}},
	{"stone", Oklch{0.97, 0.02, 50}, Oklch{0.14, 0.03, 40}},
}

func randomOKLCH() Oklch {
	palette := tailwindPalette[rand.Intn(len(tailwindPalette))]
	isDark := rand.Float64() < 0.5

	if isDark {
		// 800, 900, 950
		variant := rand.Float64()
		if variant < 0.33 {
			return Oklch{palette.dark.L - 0.08, palette.dark.C * 1.5, palette.dark.H}
		} else if variant < 0.66 {
			return palette.dark
		} else {
			return Oklch{palette.dark.L - 0.05, palette.dark.C * 0.7, palette.dark.H}
		}
	} else {
		// 50, 100, 200
		variant := rand.Float64()
		if variant < 0.33 {
			return Oklch{palette.light.L, palette.light.C * 0.5, palette.light.H}
		} else if variant < 0.66 {
			return palette.light
		} else {
			return Oklch{palette.light.L - 0.1, palette.light.C * 2, palette.light.H}
		}
	}
}

func oklchToRGB(ok Oklch) color.RGBA {
	L := ok.L
	a := ok.C * math.Cos(ok.H*math.Pi/180)
	b := ok.C * math.Sin(ok.H*math.Pi/180)

	l_ := L + 0.3963377774*a + 0.2158037573*b
	m_ := L - 0.1055613458*a - 0.0638541728*b
	s_ := L - 0.0894841775*a - 1.2914855480*b

	l := l_ * l_ * l_
	m := m_ * m_ * m_
	s := s_ * s_ * s_

	R := +4.0767416621*l - 3.3077115913*m + 0.2309699292*s
	G := -1.2684380046*l + 2.6097574011*m - 0.3413193965*s
	B := -0.0041960863*l - 0.7034186147*m + 1.7076147010*s

	clamp := func(x float64) uint8 {
		if x < 0 {
			x = 0
		}
		if x > 1 {
			x = 1
		}
		return uint8(x*255 + 0.5)
	}
	return color.RGBA{clamp(R), clamp(G), clamp(B), 255}
}

func initColors(colors []Oklch, tipoFondo int) {
	var Lfondo float64
	if tipoFondo == 0 {
		Lfondo = colors[0].L
	} else {
		Lfondo = (colors[0].L + colors[1].L) / 2
	}

	bgColors = colors

	if Lfondo < 0.4 {
		titleColor = Oklch{L: 0.92 + rand.Float64()*0.08, C: 0.08, H: rand.Float64() * 360}
		authorColor = Oklch{L: 0.75 + rand.Float64()*0.15, C: 0.08, H: rand.Float64() * 360}
		formColor = Oklch{L: 0.4 + rand.Float64()*0.3, C: 0.2 + rand.Float64()*0.2, H: rand.Float64() * 360}
	} else {
		titleColor = Oklch{L: 0.08 + rand.Float64()*0.12, C: 0.08, H: rand.Float64() * 360}
		authorColor = Oklch{L: 0.18 + rand.Float64()*0.15, C: 0.08, H: rand.Float64() * 360}
		formColor = Oklch{L: 0.4 + rand.Float64()*0.3, C: 0.2 + rand.Float64()*0.2, H: rand.Float64() * 360}
	}
}

func getFormColor() Oklch {
	return formColor
}

func getBgColors() []Oklch {
	return bgColors
}
