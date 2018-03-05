package json

import (
	"encoding/json"
	"fmt"

	"github.com/mantyr/schema"
	"github.com/mantyr/schema/constraint"
)

// Value это json реализация интерфейса schema.Value
type Value struct {
	// order содержит в себе порядковое значение
	order int64

	// name это название ключа по которому был получен объект
	name string

	// required означает что поле обязательное для родительского объекта
	required bool

	data value
}

// value это описание значения
type value struct {
	Title *string `json:"title,omitempty"`

	// Description это описание значения
	Description string `json:"description"`

	// Type это тип значения
	Type schema.Type `json:"type"`

	// Enum это перечень возможных значений
	Enum schema.Enum `json:"enum"`

	// Minimum это минимальное значение для числа
	Minimum int `json:"minimum"`

	// Maximum это максимальное значение для числа
	Maximum int `json:"maximum"`

	// MinLength это минимальная длина значения
	// применительно если тип значения string
	MinLength int `json:"minLength"`

	// MaxLength это максимальная длина значения
	// применительно если тип значения string
	MaxLength int `json:"maxLength"`

	// Properties это перечень параметров
	// применительно если значение является объектом
	Properties Properties `json:"properties"`

	// Required это перечень обязательных полей
	// применительно если значение является объектом
	Required []string `json:"required"`

	// Examples это перечень примеров для значения
	Examples schema.Examples `json:"examples"`

	// Items это описание объекта в массиве значений
	// применительно если тип значения массив
	Items *Value `json:"items"`

	// MinItems это минимальное количество значений в массиве
	MinItems int `json:"minItems"`

	// MaxItems это максимальное количество значений в массиве
	MaxItems int `json:"maxItems"`
}

// Name возвращает название значения
func (v *Value) Name() string {
	return v.name
}

// Title возвращает заголовок значения
func (v *Value) Title() string {
	if v.data.Title != nil {
		return *v.data.Title
	}
	return ""
}

// Description возвращает описание параметра
func (v *Value) Description() string {
	return v.data.Description
}

// Type возвращает тип параметра
func (v *Value) Type() schema.Type {
	return v.data.Type
}

// Values возвращает параметры объекта
func (v *Value) Values() []schema.Value {
	return v.data.Properties
}

// Constraints возвращает список ограничений
// nolint:gocyclo
func (v *Value) Constraints() (*schema.Constraints, error) {
	var constraints schema.Constraints
	switch v.Type() {
	case schema.Integer: // nolint:dupl
		if v.data.Minimum != 0 || v.data.Maximum != 0 {
			c, err := constraint.NewMinimum(v.data.Minimum, v.data.Maximum)
			if err != nil {
				return nil, err
			}
			constraints = append(constraints, c)
		}
	case schema.String: // nolint:dupl
		if v.data.MinLength != 0 || v.data.MaxLength != 0 {
			c, err := constraint.NewLength(v.data.MinLength, v.data.MaxLength)
			if err != nil {
				return nil, err
			}
			constraints = append(constraints, c)
		}
	case schema.Object:
	case schema.Boolean:
	case schema.Array:
		c, err := constraint.NewCount(v.data.MinItems, v.data.MaxItems)
		if err != nil {
			return nil, err
		}
		constraints = append(constraints, c)
	}
	if len(v.data.Enum) > 0 {
		c, err := constraint.NewEnum(v.data.Enum)
		if err != nil {
			return nil, err
		}
		constraints = append(constraints, c)
	}
	if v.required {
		c, err := constraint.NewRequired(v.required)
		if err != nil {
			return nil, err
		}
		constraints = append(constraints, c)
	}
	return &constraints, nil
}

// Examples возвращает список примеров для значения
func (v *Value) Examples() (*schema.Examples, error) {
	return &v.data.Examples, nil
}

// UnmarshalJSON необходим для декодирования JSON
func (v *Value) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &v.data)
	if err != nil {
		return err
	}
	for _, value := range v.Values() {
		if v.IsRequired(value.Name()) {
			value.SetRequired(true)
		}
	}
	return nil
}

// IsRequired возвращает true если значение обязательное
func (v *Value) IsRequired(expected string) bool {
	for _, name := range v.data.Required {
		if name == expected {
			return true
		}
	}
	return false
}

// SetRequired устанавливает флаг обязательности
func (v *Value) SetRequired(value bool) {
	v.required = value
}

// Schema возвращает схему данных на основе значения
func (v *Value) Schema() schema.Schema {
	return &Schema{
		data: *v,
	}
}

// Schemas возвращает вложенные схемы данных если таковые имеются
func (v *Value) Schemas() []schema.Schema {
	list := []schema.Schema{}
	if v.Type() == schema.Object {
		list = append(list, v.Schema())
		for _, value := range v.Values() {
			s := value.Schema()
			s.SetTitle(
				fmt.Sprintf(
					"%s -> %s",
					v.Title(),
					value.Name(),
				))
			list = append(list, s.Value().Schemas()...)
		}
	}
	return list
}
