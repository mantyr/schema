// Package sphinx это пакет для генерации Sphinx документации на основании Schema
package sphinx

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mantyr/schema"
)

// Encoder это кодировщик для sphinx-doc
type Encoder struct {
	buf  *bytes.Buffer
	line rune
}

type ParamEncoder struct {
	buf *bytes.Buffer
}

// NewEncoder возвращает новый Encoder
func NewEncoder() *Encoder {
	e := &Encoder{
		buf:  bytes.NewBuffer(make([]byte, 1000000)),
		line: '-',
	}
	e.buf.Reset()
	return e
}

// EncodeSchema кодирует схему данных в Sphinx-doc формат
// Все вложенные схемы раскрываются отдельными схемами
func (e *Encoder) Encode(s schema.Schema) error {
	if s == nil {
		return errors.New("empty schema")
	}
	title := s.Title()
	e.buf.WriteString(
		fmt.Sprintf(
			"\n%s\n%s\n\n%s\n\n",
			title,
			strings.Repeat(
				string(e.line),
				utf8.RuneCountInString(title),
			),
			s.Description(),
		),
	)
	e.line = '^'
	var err error
	value := s.Value()
	switch value.Type() {
	case schema.Object:
		for _, v := range value.Values() {
			err = e.encodeValue(v)
			if err != nil {
				return err
			}
		}
	default:
		err = e.encodeValue(value)
		if err != nil {
			return err
		}
	}
	e.buf.WriteString("\n")
	return nil
}

func (e *Encoder) encodeValue(v schema.Value) error {
	if v == nil {
		return errors.New("empty schema value")
	}
	e.buf.WriteString(
		fmt.Sprintf(
			"- ``%s``, %s: %s\n",
			v.Name(),
			v.Type(),
			v.Description(),
		),
	)
	constraints, err := v.Constraints()
	if err != nil {
		return err
	}
	if constraints == nil {
		return errors.New("unexpected nil in constraints")
	}
	for _, constraint := range *constraints {
		e.buf.WriteString(
			fmt.Sprintf(
				"	- ``%s``: %s\n",
				constraint.Name(),
				constraint.Value(),
			),
		)
	}
	return nil
}

// Bytes возвращает бинарное представление
func (e *Encoder) Bytes() []byte {
	return e.buf.Bytes()
}

// Marshal кодирует Schema в Sphinx-doc
func Marshal(s []schema.Schema) ([]byte, error) {
	var err error
	e := NewEncoder()
	for _, schema := range s {
		err = e.Encode(schema)
		if err != nil {
			return []byte{}, err
		}
	}
	return e.Bytes(), nil
}
