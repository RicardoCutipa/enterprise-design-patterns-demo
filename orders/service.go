package orders

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) PlaceOrder(order Order) (Order, error) {
	order.Total = order.CalculateTotal()
	if order.Total <= 0 {
		return Order{}, fmt.Errorf("invalid order total")
	}

	uow := NewUnitOfWork(s.repo)
	uow.Register(order)

	if err := uow.Commit(); err != nil {
		return Order{}, err
	}

	return order, nil
}
