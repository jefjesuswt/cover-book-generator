package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jefjesuswt/cover-book-generator/internal/cover"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // ponytail: .env optional, ignore error if missing

	var title, author, output string
	flag.StringVar(&title, "title", "", "Título del libro")
	flag.StringVar(&author, "author", "", "author del libro")
	flag.StringVar(&output, "output", "miniatura.png", "Archivo de output")
	flag.Parse()

	if title == "" || author == "" {
		log.Fatal("Debes proporcionar título y author")
	}

	apiKey := os.Getenv("UNSPLASH_ACCESS_KEY")

	err := cover.Generate(title, author, output, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Miniatura generada: %s\n", output)
}
