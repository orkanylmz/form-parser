package generator

import (
	"errors"
	"form-parser/domain"
)

// NewGenerator returns the correct generator depending on the file extension
func NewGenerator(fileType string) (domain.Generator, error) {
	switch fileType {
	case "pdf":
		return &PDFGenerator{}, nil
	default:
		return nil, errors.New("Unsupported file type: " + fileType)
	}
}
