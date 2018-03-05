package schema

// Enum это перечень возможных значений
type Enum []EnumValue

// EnumValue это возможное значение
type EnumValue string

// UnmarshalJSON необходим для разбора значений неопределённого типа
func (v *EnumValue) UnmarshalJSON(data []byte) error {
	*v = EnumValue(string(data))
	return nil
}

// Values возвращает список возможных значений
func (e Enum) Values() []string {
	values := make([]string, len(e))
	for i, value := range e {
		values[i] = string(value)
	}
	return values
}
