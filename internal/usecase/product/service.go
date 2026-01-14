package product

import (
	"errors"

	"github.com/imbafff/product-warehouse-api/internal/entity"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(p *entity.Product) (int64, error) {
	if p.Name == "" {
		return 0, errors.New("name is required")
	}
	if p.Price <= 0 {
		return 0, errors.New("price must be greater than zero")
	}
	if p.Quantity < 0 {
		return 0, errors.New("quantity must be non-negative")
	}

	return s.repo.Create(p)
}

func (s *Service) GetByID(id int64) (*entity.Product, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.repo.GetByID(id)
}

func (s *Service) Update(id int64, p *entity.Product) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Quantity < 0 {
		return errors.New("quantity must be non-negative")
	}

	return s.repo.Update(id, p)
}

func (s *Service) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.repo.Delete(id)
}

func (s *Service) GetAll() ([]*entity.Product, error) {
	return s.repo.GetAll()
}
