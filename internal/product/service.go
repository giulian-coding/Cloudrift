package product

import "context"

type Repository interface {
	Create(context.Context, Product) error
	ByID(context.Context, string) (Product, error)
}
