package parser

import (
	"form-parser/domain"
)

type JSONParser struct{}

func (x *JSONParser) Parse(data []byte) (domain.Form, error) {

	return domain.Form{}, nil
}
