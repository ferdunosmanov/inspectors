package GetAllRegistrations

import "github.com/ferdunosmanov/inspectors/pkg/adding"

type Repository interface {
	GetAllRegistrations() ([]adding.Registrations, error)
}

type Service interface {
	GetAllRegistrations() ([]adding.Registrations, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAllRegistrations() ([]adding.Registrations, error) {
	ir, err := s.r.GetAllRegistrations()
	if err != nil {
		return nil, err
	}
	return ir, nil
}
