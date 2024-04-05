package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
	List(pageSize int, page int) ([]Order, error)
}
