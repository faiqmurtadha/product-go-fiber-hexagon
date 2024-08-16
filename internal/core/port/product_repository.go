package port

import "product-go-fiber-hexagon/internal/core/model"

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	GetAll(page, limit int64, name string) (*model.PaginatedProduct, error)
	FindById(id string) (*model.Product, error)
	Update(id string, product *model.Product) (*model.Product, error)
	Delete(id string) error
}
