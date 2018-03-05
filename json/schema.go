package json

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mantyr/schema"
)

// Schema это json реализация схемы данных
type Schema struct {
	data Value
}

// Title возвращает заголовок схемы данных
func (s *Schema) Title() string {
	if s.data.data.Title != nil {
		return *s.data.data.Title
	}
	return ""
}

// SetTitle устанавливает заголовок схемы данных
func (s *Schema) SetTitle(title string) {
	s.data.data.Title = &title
}

// Description возвращает описание схемы данных
func (s *Schema) Description() string {
	return s.data.data.Description
}

// Type возвращает тип значения схемы данных
func (s *Schema) Type() schema.Type {
	return schema.Type(strings.ToLower(string(s.data.data.Type)))
}

// Value возвращает значение схемы данных
func (s *Schema) Value() schema.Value {
	return &s.data
}

// Schemas возвращает все вложенные схемы данных
func (s *Schema) Schemas() []schema.Schema {
	return s.data.Schemas()
}

// NewSchemas возвращает новый список схем данных
func NewSchemas(data []byte) ([]schema.Schema, error) {
	s := &Schema{}
	err := json.Unmarshal(data, s.Value())
	if err != nil {
		return nil, err
	}
	list := s.Schemas()
	return list, nil
}

// NewFileSchemas возвращает новый список схем данных на основе файла
func NewFileSchemas(address string) ([]schema.Schema, error) {
	file, err := os.Open(address)
	if err != nil {
		return nil, err
	}
	defer file.Close() // nolint:errcheck

	return NewReaderSchemas(file)
}

// NewReaderSchemas возвращает новый список схем данных на основе io.Reader
func NewReaderSchemas(r io.Reader) ([]schema.Schema, error) {
	if r == nil {
		return nil, errors.New("empty io.Reader")
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewSchemas(data)
}
