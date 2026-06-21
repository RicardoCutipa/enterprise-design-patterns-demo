package orders

type Record struct {
	ID    string
	Total float64
}

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

func (m Mapper) ToRecord(order Order) Record {
	return Record{
		ID:    order.ID,
		Total: order.Total,
	}
}

func (m Mapper) ToOrder(record Record) Order {
	return Order{
		ID:    record.ID,
		Total: record.Total,
	}
}
