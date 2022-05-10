package reading

type Repository interface {
	GetAllProductNames() ([]string, error)
}

type Service interface {
	GetAllProductNames() ([]string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllProductNames() ([]string, error) {
	cs, err := s.r.GetAllProductNames()
	if err != nil {
		return nil, err
	}
	return cs, nil
}
