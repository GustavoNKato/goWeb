package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(productName, color string, price float64, amount int) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetById(id int) (Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return product, err
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(productName, color string, price float64, amount int) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, productName, color, price, amount)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
