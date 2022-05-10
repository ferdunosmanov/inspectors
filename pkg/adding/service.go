package adding

type Repository interface {
	AddProduct(Product) ([]string, error)
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
