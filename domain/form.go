package domain

type Form struct {
	Fields   []Field
	Sections []Section
}

type Field struct {
	Name      string
	Type      string
	Optional  bool
	FieldType FieldType
	Caption   string
	Labels    []Label
}

type FieldType string

type Label struct {
	Name string
	Text string
}

type Section struct {
	Name     string
	Optional bool
	Title    string
	Fields   []Field
}
