package product

import "github.com/imbafff/product-warehouse-api/internal/entity"

type UseCase interface {
	Create(product *entity.Product) (int64, error)
	GetByID(id int64) (*entity.Product, error)
	Update(id int64, product *entity.Product) error
	Delete(id int64) error
	GetAll() ([]*entity.Product, error)
}
