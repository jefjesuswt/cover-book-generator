# Book Cover Generator v2

## Decisions
- Image source: Unsplash API (keyword search, random pick)
- Image usage: Full background + semi-transparent overlay for text readability
- Fonts: Bundle TTF files by theme (~8 Google Fonts)
- Theme: Auto-detect from title keywords

## Phase 1: Theme system
- [ ] Create `internal/cover/theme.go` — theme detection from title keywords
  - Themes: medieval, cyber, elegant, minimal, tech, business, classic, nature
  - Each theme maps: font file, font weight preference, overlay style
  - Keyword matching on title (e.g. "go"→tech, "economic"→business)
- [ ] Download + bundle ~8 Google Fonts TTFs into `fonts/` by theme
  - medieval: Cinzel (serif)
  - cyber: Orbitron (geometric sans)
  - elegant: Playfair Display (serif)
  - minimal: Inter (clean sans)
  - tech: JetBrains Mono (mono)
  - business: Merriweather (serif)
  - classic: Lora (transitional serif)
  - nature: Cormorant Garamond (serif)

## Phase 2: Image fetching
- [ ] Create `internal/cover/image.go` — Unsplash API client
  - `FetchImage(apiKey, query string) (image.Image, error)`
  - GET `https://api.unsplash.com/photos/random?query={query}&orientation=portrait`
  - Parse response, download `regular` size URL
  - Resize/crop to 1600x2560 (center crop)
  - Env var `UNSPLASH_ACCESS_KEY` for API key

## Phase 3: Color extraction
- [ ] Create `internal/cover/color.go` additions — extract dominant colors from image
  - Simple median-cut or k-means on downsampled image pixels
  - Return: dominant bg color, text color (light/dark based on luminance)
  - Used for overlay tint + text color selection

## Phase 4: Generator rewrite
- [ ] Modify `internal/cover/generator.go` — new pipeline
  1. `detectTheme(title)` → theme config
  2. `fetchImage(apiKey, themeQuery)` → background image
  3. `extractColors(img)` → overlay color + text color
  4. Draw image as full background
  5. Draw semi-transparent overlay (gradient: stronger at bottom for text)
  6. Draw text with theme font (title + author)
- [ ] Modify `internal/cover/text.go` — accept theme, use theme font
  - Keep auto-sizing logic (600→60)
  - Load font from theme config instead of hardcoded Roboto
- [ ] Remove or gut `internal/cover/shapes.go` — no longer needed with image backgrounds

## Phase 5: Wire up
- [ ] Modify `main.go` — read `UNSPLASH_ACCESS_KEY` from env, pass to generator
- [ ] Test with each of the 9 existing books
