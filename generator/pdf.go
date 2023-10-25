package generator

import (
	"bytes"
	"form-parser/domain"
	"github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct{}

func (pg *PDFGenerator) Generate(form domain.Form) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Add a title
	pdf.Cell(190, 10, "Form Details")
	pdf.Ln(20)

	// Render fields outside of sections first
	for _, field := range form.Fields {
		renderField(pdf, field)
		pdf.Ln(10)
	}

	// Render sections
	for _, section := range form.Sections {
		pdf.SetFont("Arial", "BU", 14)
		pdf.Cell(190, 10, section.Title)
		pdf.Ln(20)

		for _, field := range section.Fields {
			renderField(pdf, field)
			pdf.Ln(10)
		}
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func renderField(pdf *gofpdf.Fpdf, field domain.Field) {
	switch field.FieldType {
	case "Select":
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(190, 10, field.Caption)
		pdf.Ln(10)
		for _, label := range field.Labels {
			pdf.SetFont("Arial", "", 12)
			pdf.Cell(190, 10, "- "+label.Text)
			pdf.Ln(10)
		}
	case "TextBox":
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(190, 10, field.Caption)
		pdf.Ln(15)
		x, y := pdf.GetXY()
		drawTextInputBox(pdf, x, y, 90, 10)
		pdf.Ln(15)

	case "File":
		pdf.Cell(190, 10, field.Caption)
		pdf.Ln(20)
	}

}

func drawTextInputBox(pdf *gofpdf.Fpdf, x float64, y float64, width float64, height float64) {
	pdf.SetFillColor(255, 255, 255)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Rect(x, y, width, height, "FD")
}
