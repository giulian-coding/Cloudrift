package order

import (
	"context"

	"time"

	"uuid"
)

type Repository interface {
	Create(context.Context, Order) error
	ByID(context.Context, string) (Order, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (Order, error) {
	if req.CustomerID == "" {
		return Order{}, ErrCustomerRequired
	}

	newID := uuid.New().String()
	order := Order{
		ID:         newID,
		CustomerID: req.CustomerID,
		Status:     StatusPending,
		CreatedAt:  time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return Order{}, err
	}

	return order, nil
}

func (s *Service) Get(ctx context.Context, id string) (Order, error) {
	order, err := s.repo.ByID(ctx, id)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}
