package orders

import "fmt"

type Repository interface {
	Save(order Order) error
	FindByID(id string) (Order, error)
}

type MemoryRepository struct {
	storage map[string]Record
	mapper  Mapper
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		storage: map[string]Record{},
		mapper:  NewMapper(),
	}
}

func (r *MemoryRepository) Save(order Order) error {
	record := r.mapper.ToRecord(order)
	r.storage[record.ID] = record
	return nil
}

func (r *MemoryRepository) FindByID(id string) (Order, error) {
	record, ok := r.storage[id]
	if !ok {
		return Order{}, fmt.Errorf("order %s not found", id)
	}

	return r.mapper.ToOrder(record), nil
}
