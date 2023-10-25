package domain

// FormRepository defines the interface that all parsers should implement
type FormRepository interface {
	Parse(data []byte) (Form, error)
}
