package service

import (
	"product-go-fiber-hexagon/internal/core/model"
	"product-go-fiber-hexagon/internal/core/port"
)

type ProductService struct {
	repository port.ProductRepository
}

func NewProductService(repository port.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (srvc *ProductService) Create(product *model.Product) (*model.Product, error) {
	return srvc.repository.Create(product)
}

func (srvc *ProductService) GetAll(page, limit int64, name string) (*model.PaginatedProduct, error) {
	return srvc.repository.GetAll(page, limit, name)
}

func (srvc *ProductService) FindById(id string) (*model.Product, error) {
	return srvc.repository.FindById(id)
}

func (srvc *ProductService) Update(id string, product *model.Product) (*model.Product, error) {
	return srvc.repository.Update(id, product)
}

func (srvc *ProductService) Delete(id string) error {
	return srvc.repository.Delete(id)
}
