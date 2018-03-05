package constraint // nolint:dupl

import (
	"fmt"
)

// Minimum это огранчиение по количеству элементов
type Minimum struct {
	min int
	max int
}

// NewMinimum возвращает новое ограничение по длине строки
func NewMinimum(min, max int) (*Minimum, error) {
	if min > max {
		return nil, fmt.Errorf("expected min <= max but actual %d > %d", min, max)
	}
	return &Minimum{
		min: min,
		max: max,
	}, nil
}

// Name возвращает название ограничения
func (l *Minimum) Name() string {
	return "значение"
}

// Value возвращает значение ограничения
// nolint:dupl
func (l *Minimum) Value() string {
	switch {
	case l.max > 0:
		return fmt.Sprintf("от %d до %d", l.min, l.max)
	case l.min > 0:
		return fmt.Sprintf("от %d", l.min)
	}
	return ""
}
