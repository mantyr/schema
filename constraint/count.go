package constraint // nolint:dupl

import (
	"fmt"
)

// Count это огранчиение по количеству элементов
type Count struct {
	min int
	max int
}

// NewCount возвращает новое ограничение по длине строки
func NewCount(min, max int) (*Count, error) {
	if min > max {
		return nil, fmt.Errorf("expected min <= max but actual %d > %d", min, max)
	}
	if min < 0 {
		return nil, fmt.Errorf("expected min сount >= 0 but actual %d", min)
	}
	if max < 0 {
		return nil, fmt.Errorf("expected max count >= 0 but actual %d", max)
	}
	return &Count{
		min: min,
		max: max,
	}, nil
}

// Name возвращает название ограничения
func (l *Count) Name() string {
	return "количество элементов"
}

// Value возвращает значение ограничения
// nolint:dupl
func (l *Count) Value() string {
	switch {
	case l.max > 0:
		return fmt.Sprintf("от %d до %d", l.min, l.max)
	case l.min > 0:
		return fmt.Sprintf("от %d", l.min)
	}
	return ""
}
