package stocks

import (
	"encoding/json"
	"fmt"
)

type JsonParser[T any] struct{}

func (jp *JsonParser[T]) ParseResponseBody(body []byte) (*T, error) {
	var result T
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return &result, nil
}
