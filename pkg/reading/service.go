package reading

import "github.com/ferdunosmanov/inspectors/pkg/adding"

type Repository interface {
	GetAllProductNames() ([]adding.Product, error)
}

type Service interface {
	GetAllProductNames() ([]adding.Product, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllProductNames() ([]adding.Product, error) {
	cs, err := s.r.GetAllProductNames()
	if err != nil {
		return nil, err
	}
	return cs, nil
}
