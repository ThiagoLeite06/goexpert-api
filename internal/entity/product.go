package entity

import (
	"github.com/ThiagoLeite06/goexpert/api-9/pkg/entity"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrorIDIsRequired    = errors.New("ID is required")
	ErrorInvalidID       = errors.New("Invalid ID")
	ErrorNameIsRequired  = errors.New("Name is required!")
	ErrorPriceIsRequired = errors.New("Price is required!")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}

	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil

}
