package entity

import (
	"context"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id int) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id int) error
}
