package cover

import (
	"strings"
)

type Theme struct {
	Name       string
	FontFile   string
	Query      string  // Unsplash search query
	Overlay    float64 // overlay opacity 0-1
}

var themes = map[string]Theme{
	"medieval": {
		Name:     "medieval",
		FontFile: "fonts/Cinzel.ttf",
		Query:    "medieval architecture gothic",
		Overlay:  0.50,
	},
	"cyber": {
		Name:     "cyber",
		FontFile: "fonts/Orbitron.ttf",
		Query:    "cyberpunk neon technology",
		Overlay:  0.65,
	},
	"elegant": {
		Name:     "elegant",
		FontFile: "fonts/PlayfairDisplay.ttf",
		Query:    "elegant abstract art",
		Overlay:  0.45,
	},
	"minimal": {
		Name:     "minimal",
		FontFile: "fonts/Inter.ttf",
		Query:    "minimalist abstract",
		Overlay:  0.50,
	},
	"tech": {
		Name:     "tech",
		FontFile: "fonts/JetBrainsMono.ttf",
		Query:    "technology programming code",
		Overlay:  0.60,
	},
	"business": {
		Name:     "business",
		FontFile: "fonts/Merriweather.ttf",
		Query:    "business economics finance",
		Overlay:  0.55,
	},
	"classic": {
		Name:     "classic",
		FontFile: "fonts/Lora.ttf",
		Query:    "classic library vintage books",
		Overlay:  0.50,
	},
	"nature": {
		Name:     "nature",
		FontFile: "fonts/CormorantGaramond.ttf",
		Query:    "nature landscape forest",
		Overlay:  0.50,
	},
}

// keyword → theme name
var keywordMap = map[string]string{
	// medieval
	"cathedral": "medieval", "bazaar": "medieval", "castle": "medieval",
	"kingdom": "medieval", "medieval": "medieval", "fantasy": "medieval",
	"dragon": "medieval", "lord": "medieval", "sword": "medieval",
	// cyber
	"cyber": "cyber", "neural": "cyber", "ai": "cyber", "digital": "cyber",
	"machine": "cyber", "robot": "cyber", "hacker": "cyber", "matrix": "cyber",
	// elegant
	"art": "elegant", "design": "elegant", "beauty": "elegant",
	"aesthetic": "elegant", "elegance": "elegant", "poetry": "elegant",
	// minimal
	"minimal": "minimal", "simple": "minimal", "clean": "minimal",
	"essential": "minimal",
	// tech
	"code": "tech", "system": "tech", "backend": "tech", "programming": "tech",
	"software": "tech", "typescript": "tech", "zig": "tech", "bun": "tech",
	"elysia": "tech", "computer": "tech", "turing": "tech", "algorithm": "tech",
	"data": "tech", "developer": "tech", "engineering": "tech",
	"performance": "tech", "go": "tech", "systemantics": "tech",
	// business
	"economic": "business", "economy": "business", "market": "business",
	"business": "business", "philosophy": "business", "social": "business",
	"political": "business", "wealth": "business", "capital": "business",
	"trade": "business", "rawls": "business", "nozick": "business",
	"rothbard": "business", "pensamiento": "business",
	// classic
	"history": "classic", "thought": "classic", "vision": "classic",
	"science": "classic", "wisdom": "classic", "knowledge": "classic",
	"mind": "classic",
	// nature
	"nature": "nature", "forest": "nature", "earth": "nature",
	"green": "nature", "environment": "nature", "garden": "nature",
	"wild": "nature", "tree": "nature",
}

func detectTheme(title string) Theme {
	words := strings.Fields(strings.ToLower(title))
	for _, w := range words {
		w = strings.Trim(w, ".,;:!?\"'()[]{}—–-")
		if themeName, ok := keywordMap[w]; ok {
			return themes[themeName]
		}
	}
	return themes["minimal"]
}
