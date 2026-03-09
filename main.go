package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jefjesuswt/cover-book-generator/internal/cover"
)

func main() {
	var title, author, output string
	flag.StringVar(&title, "title", "", "Título del libro")
	flag.StringVar(&author, "author", "", "author del libro")
	flag.StringVar(&output, "output", "miniatura.png", "Archivo de output")
	flag.Parse()

	if title == "" || author == "" {
		log.Fatal("Debes proporcionar título y author")
	}

	err := cover.Generate(title, author, output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Miniatura generada: %s\n", output)
}
