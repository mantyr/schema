package constraint // nolint:dupl

import (
	"errors"
	"strings"

	"github.com/mantyr/schema"
)

// Enum это огранчиение по количеству элементов
type Enum struct {
	data []string
}

// NewEnum возвращает новое ограничение по длине строки
func NewEnum(data schema.Enum) (*Enum, error) {
	if len(data) < 0 {
		return nil, errors.New("empty enum")
	}
	e := &Enum{
		data: make([]string, len(data)),
	}
	for i, v := range data {
		e.data[i] = string(v)
	}
	return e, nil
}

// Name возвращает название ограничения
func (e *Enum) Name() string {
	return "возможные значения"
}

// Value возвращает значение ограничения
func (e *Enum) Value() string {
	return strings.Join(e.data, ", ")
}
