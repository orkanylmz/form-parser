package parser

import (
	"encoding/xml"
	"form-parser/domain"
	"strings"
)

type xmlForm struct {
	Fields   []xmlField   `xml:"Field"`
	Sections []xmlSection `xml:"Section"`
}

func (x xmlForm) toDomain() (domain.Form, error) {
	form := domain.Form{}

	for _, xmlF := range x.Fields {
		field, err := xmlF.toDomain()
		if err != nil {
			return domain.Form{}, err
		}
		form.Fields = append(form.Fields, field)
	}

	for _, xmlSec := range x.Sections {
		section, err := xmlSec.toDomain()
		if err != nil {
			return domain.Form{}, nil
		}
		form.Sections = append(form.Sections, section)
	}
	return form, nil
}

type xmlField struct {
	Name      string     `xml:"Name,attr"`
	Type      string     `xml:"Type,attr"`
	Optional  string     `xml:"Optional,attr"`
	FieldType string     `xml:"FieldType,attr"`
	Caption   string     `xml:"Caption"`
	Labels    []xmlLabel `xml:"Labels>Label"`
}

func (f xmlField) toDomain() (domain.Field, error) {
	var labels []domain.Label
	for _, l := range f.Labels {
		labels = append(labels, domain.Label{
			Name: l.Name,
			Text: l.Text,
		})
	}

	return domain.Field{
		Name:      f.Name,
		Type:      f.Type,
		Optional:  strings.ToLower(f.Optional) == "true",
		FieldType: domain.FieldType(f.FieldType),
		Caption:   f.Caption,
		Labels:    labels,
	}, nil
}

type xmlLabel struct {
	Name string `xml:"Name,attr"`
	Text string `xml:",chardata"`
}

type xmlSection struct {
	Name     string     `xml:"Name,attr"`
	Optional string     `xml:"Optional,attr"`
	Title    string     `xml:"Title"`
	Fields   []xmlField `xml:"Contents>Field"`
}

func (s xmlSection) toDomain() (domain.Section, error) {
	var fields []domain.Field

	for _, xmlF := range s.Fields {
		field, err := xmlF.toDomain()
		if err != nil {
			return domain.Section{}, err
		}
		fields = append(fields, field)
	}

	return domain.Section{
		Name:     s.Name,
		Optional: strings.ToLower(s.Optional) == "true",
		Title:    s.Title,
		Fields:   fields,
	}, nil
}

type XMLParser struct{}

func (x *XMLParser) Parse(data []byte) (domain.Form, error) {
	xmlF := xmlForm{}
	err := xml.Unmarshal(data, &xmlF)
	if err != nil {
		return domain.Form{}, err
	}

	form, err := xmlF.toDomain()
	if err != nil {
		return domain.Form{}, err
	}

	return form, nil
}
