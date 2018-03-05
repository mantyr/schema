package json

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/mantyr/schema"
)

// Properties это список полей объекта
type Properties []schema.Value

// UnmarshalJSON необходим для сохранения порядка полей в объекте
func (p *Properties) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(strings.NewReader(string(data)))
	order := int64(0)
	for {
		token, err := dec.Token()
		if err == io.EOF {
			return nil
		}
		switch name := token.(type) {
		case json.Delim:
		case string:
			v := &Value{
				order: order,
				name:  name,
			}
			order++
			err = dec.Decode(&v.data)
			if err != nil {
				return err
			}
			(*p) = append(*p, v)
		}
	}
}
