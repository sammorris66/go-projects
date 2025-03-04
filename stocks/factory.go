package stocks

import "errors"

type Factory struct {
	registry map[string]func(string) (BaseExchange, error)
}

func NewFactory() *Factory {
	return &Factory{
		registry: make(map[string]func(string) (BaseExchange, error)),
	}
}

func (f *Factory) Register(exchangeType string, constructor func(string) (BaseExchange, error)) {
	f.registry[exchangeType] = constructor
}

func (f *Factory) Create(exchangeType, symbol string) (BaseExchange, error) {
	if constructor, found := f.registry[exchangeType]; found {
		return constructor(symbol)
	}
	return nil, errors.New("unknown exchange type")
}

var GlobalFactory = NewFactory() // Global factory instance
