package domain

type FormService struct {
	repo FormRepository
	gen  Generator
}

func NewFormService(repo FormRepository, gen Generator) *FormService {
	return &FormService{repo: repo, gen: gen}
}

func (s *FormService) Generate(data []byte) ([]byte, error) {
	form, err := s.repo.Parse(data)

	if err != nil {
		return nil, err
	}

	return s.gen.Generate(form)
}
