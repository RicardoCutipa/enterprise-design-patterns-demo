package orders

import "testing"

func TestPlaceOrderSavesOrder(t *testing.T) {
	repo := NewMemoryRepository()
	service := NewService(repo)

	order := Order{
		ID: "ORD-2001",
		Items: []Item{
			{Name: "Keyboard", Price: 50, Quantity: 1},
			{Name: "Mouse", Price: 25, Quantity: 2},
		},
	}

	savedOrder, err := service.PlaceOrder(order)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if savedOrder.Total != 100 {
		t.Fatalf("expected total 100, got %.2f", savedOrder.Total)
	}

	storedOrder, err := repo.FindByID("ORD-2001")
	if err != nil {
		t.Fatalf("expected saved order: %v", err)
	}

	if storedOrder.Total != 100 {
		t.Fatalf("expected stored total 100, got %.2f", storedOrder.Total)
	}
}

func TestPlaceOrderRejectsEmptyOrder(t *testing.T) {
	repo := NewMemoryRepository()
	service := NewService(repo)

	_, err := service.PlaceOrder(Order{ID: "ORD-2002"})
	if err == nil {
		t.Fatal("expected error")
	}
}
