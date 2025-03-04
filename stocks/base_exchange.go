package stocks

type BaseExchange interface {
	GetPrice() (float64, error)
}
