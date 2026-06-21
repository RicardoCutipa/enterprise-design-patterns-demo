package orders

type UnitOfWork struct {
	pending []Order
	repo    Repository
}

func NewUnitOfWork(repo Repository) *UnitOfWork {
	return &UnitOfWork{repo: repo}
}

func (u *UnitOfWork) Register(order Order) {
	u.pending = append(u.pending, order)
}

func (u *UnitOfWork) Commit() error {
	for _, order := range u.pending {
		if err := u.repo.Save(order); err != nil {
			return err
		}
	}

	u.pending = nil
	return nil
}
