package schema

// Перечень типов данных в схеме данных
const (
	Integer Type = "integer"
	String  Type = "string"
	Boolean Type = "boolean"
	Object  Type = "object"
	Array   Type = "array"
)

// Type это тип данных в схеме данных
type Type string

// Schema это интерфейс схемы данных,
// каждый вложенный объект может быть представлен отдельной схемой
type Schema interface {
	// Title возвращает заголовок схемы данных
	Title() string

	// SetTitle устанавливает заголовок схемы данных
	SetTitle(title string)

	// Description возвращает описание схемы данных
	Description() string

	// Type возвращает тип схемы данных
	Type() Type

	// Value возвращает интерфейс значения
	Value() Value
}

// Value это интерфейс значения
type Value interface {
	// Name  возвращает название параметра
	Name() string

	// Description возвращает описание параметра
	Description() string

	// Type возвращает тип параметра
	Type() Type

	// Values возвращает список параметров если значение объект
	Values() []Value

	// Contraints возвращает список ограничений
	Constraints() (*Constraints, error)

	// Examples возвращает список примеров для значения
	Examples() (*Examples, error)

	// Schema возвращает схему данных на основе значения
	Schema() Schema

	// Schemas возвращает вложенные схемы данных если таковые имеются
	Schemas() []Schema

	// SetRequired устанавливает флаг обязательности
	SetRequired(value bool)
}

// Constraints это список ограничений для значения
type Constraints []Constraint

// Constraint это интерфейс ограничения
type Constraint interface {
	// Name возвращает название ограничения
	Name() string

	// Value возвращает значение ограничения
	Value() string
}

// Examples это список примеров для значения
type Examples []string
