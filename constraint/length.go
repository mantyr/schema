package constraint // nolint:dupl

import (
	"fmt"
)

// Length это огранчиение по длине
type Length struct {
	min int
	max int
}

// NewLength возвращает новое ограничение по длине строки
func NewLength(min, max int) (*Length, error) {
	if min > max {
		return nil, fmt.Errorf("expected min <= max but actual %d > %d", min, max)
	}
	if min < 0 {
		return nil, fmt.Errorf("expected min length >= 0 but actual %d", min)
	}
	if max < 0 {
		return nil, fmt.Errorf("expected max length >= 0 but actual %d", max)
	}
	return &Length{
		min: min,
		max: max,
	}, nil
}

// Name возвращает название ограничения
func (l *Length) Name() string {
	return "длина"
}

// Value возвращает значение ограничения
// nolint:dupl
func (l *Length) Value() string {
	switch {
	case l.max > 0:
		return fmt.Sprintf("от %d до %d", l.min, l.max)
	case l.min > 0:
		return fmt.Sprintf("от %d", l.min)
	}
	return ""
}
