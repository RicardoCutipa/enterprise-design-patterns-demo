package orders

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	ID    string
	Items []Item
	Total float64
}

func (o Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
