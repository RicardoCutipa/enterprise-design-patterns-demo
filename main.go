package main

import (
	"fmt"

	"github.com/RicardoCutipa/enterprise-design-patterns-demo/orders"
)

func main() {
	repository := orders.NewMemoryRepository()
	service := orders.NewService(repository)

	order := orders.Order{
		ID: "ORD-2001",
		Items: []orders.Item{
			{Name: "Keyboard", Price: 50, Quantity: 1},
			{Name: "Mouse", Price: 25, Quantity: 2},
		},
	}

	receipt, err := service.PlaceOrder(order)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("order %s saved with total $%.2f\n", receipt.ID, receipt.Total)
}
