package constraint // nolint:dupl

// Required это ограничение по обязательности поля
type Required struct {
	value bool
}

// NewRequired возвращает новое ограничение по длине строки
func NewRequired(value bool) (*Required, error) {
	return &Required{
		value: value,
	}, nil
}

// Name возвращает название ограничения
func (r *Required) Name() string {
	return "обязательный"
}

// Value возвращает значение ограничения
func (r *Required) Value() string {
	if r.value {
		return "да"
	}
	return "нет"
}
