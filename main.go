package main

import (
	"fmt"

	"github.com/RicardoCutipa/enterprise-design-patterns-demo/orders"
)

func main() {
	repository := orders.NewMemoryRepository()
	service := orders.NewService(repository)

	fmt.Println("enterprise design patterns demo in Go")
	fmt.Println("1. the domain calculates the total")
	fmt.Println("2. the service validates the order")
	fmt.Println("3. the unit of work commits the change")
	fmt.Println("4. the repository stores a mapped record")

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

	storedOrder, err := repository.FindByID(receipt.ID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("order %s saved with total $%.2f\n", receipt.ID, receipt.Total)
	fmt.Printf("stored order %s confirmed with total $%.2f\n", storedOrder.ID, storedOrder.Total)
}
