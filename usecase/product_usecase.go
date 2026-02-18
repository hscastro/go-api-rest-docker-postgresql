package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	// Repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repo}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProducts(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(idProduct int) error {

	_, err := pu.repository.GetProductById(idProduct)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("produto não encontrado")
		}
		return err
	}

	if err := pu.repository.DeleteProductById(idProduct); err != nil {
		return err
	}

	return nil
}

func (pu *ProductUsecase) UpdateteProduct(product model.Product, product_id int) (model.Product, error) {
	productId, err := pu.repository.UpdateProduct(product, product_id)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}
