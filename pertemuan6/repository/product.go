package repository

import (
	"errors"
	"pertemuan6/model"
	"time"

	"gorm.io/gorm"
)

type ProductLocalRepo struct {
	products []*model.Product
}

func (p *ProductLocalRepo) Get() ([]*model.Product, error) {
	if len(p.products) <= 0 {
		return []*model.Product{}, nil
	}
	products := []*model.Product{}
	for _, product := range p.products {
		if product.DeletedAt != nil {
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductLocalRepo) Create(product *model.Product) error {
	product.ID = uint64(len(p.products) + 1)
	p.products = append(p.products, product)
	return nil
}

func (s *ProductLocalRepo) Update(id uint64, productUpdate *model.ProductUpdate) error {
	for _, product := range s.products {
		if product.ID == id && product.DeletedAt == nil {
			product.Name = productUpdate.Name
			product.Price = productUpdate.Price
			return nil
		}
	}
	return errors.New("Product not found")
}

func (s *ProductLocalRepo) Delete(id uint64) error {
	for _, product := range s.products {
		if product.ID == id && product.DeletedAt == nil {
			tn := time.Now()
			product.DeletedAt = &tn
			return nil
		}
	}
	return errors.New("Product not found")
}

type ProductPgRepo struct {
	DB *gorm.DB
}

func (p *ProductPgRepo) Get() ([]*model.Product, error) {
	products := []*model.Product{}
	err := p.DB.Debug().Find(&products).Error
	return products, err
}

func (p *ProductPgRepo) Create(Product *model.Product) error {
	err := p.DB.Debug().Create(&Product).Error
	return err
}

func (p *ProductPgRepo) Update(id uint64, ProductUpdate *model.ProductUpdate) error {
	err := p.DB.Debug().
		Where("id = ?", id).
		Updates(&model.Product{
			Name:  ProductUpdate.Name,
			Price: ProductUpdate.Price,
		}).Error
	return err
}

func (p *ProductPgRepo) Delete(id uint64) error {
	err := p.DB.Debug().
		Where("id = ?", id).
		Delete(&model.Product{}).Error
	return err
}
