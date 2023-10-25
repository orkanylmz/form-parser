package main

import (
	"bytes"
	"form-parser/domain"
	"form-parser/generator"
	"form-parser/parser"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	inputFile := "example.xml"

	// detect the file extension, so we can load proper parser for it
	fileExt := getFileExtension(inputFile)

	data, err := readFile(inputFile)

	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// create a parser
	repo, err := parser.NewParser(fileExt)

	if err != nil {
		log.Fatalf("error detecting parser for file type: %v", err)
	}

	// create a generator
	gen, err := generator.NewGenerator("pdf")

	if err != nil {
		log.Fatalf("error creating generator for file type: %v", err)
	}

	// create the service
	service := domain.NewFormService(repo, gen)

	output, err := service.Generate(data)

	// we can export the output bytes to pdf file
	if err != nil {
		log.Fatalf("error processing form: %v", err)
	}

	outputFile, err := os.Create("output.pdf")
	if err != nil {
		log.Fatalf("error creating PDF file: %v", err)
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, bytes.NewReader(output))
	if err != nil {
		log.Fatalf("error writing to PDF file: %v", err)
	}

}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func getFileExtension(filename string) string {
	return strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
}
