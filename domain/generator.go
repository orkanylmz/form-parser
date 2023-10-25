package domain

// Generator defines an interface that all generators should implement
type Generator interface {
	Generate(form Form) ([]byte, error)
}
