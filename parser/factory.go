package parser

import (
	"errors"
	"form-parser/domain"
	"strings"
)

// NewParser returns the correct parser depending on the file extension
func NewParser(fileType string) (domain.FormRepository, error) {
	switch strings.ToLower(fileType) {
	case "xml":
		return &XMLParser{}, nil
	case "json":
		return &JSONParser{}, nil
	default:
		return nil, errors.New("Unsupported file type: " + fileType)
	}
}
