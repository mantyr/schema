package sphinx

import (
	"testing"
	"io/ioutil"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/mantyr/schema/json"
)

func TestJsonSchema(t *testing.T) {
	Convey("Проверяем кодирование в Sphinx-doc", t, func() {
		schemas, err := json.NewFileSchemas("./testdata/response.json")
		So(err, ShouldBeNil)
		So(
			len(schemas),
			ShouldEqual,
			3,
		)
		
		expected, err := ioutil.ReadFile("./testdata/response.rst")
		So(err, ShouldBeNil)

		data, err := Marshal(schemas)
		So(err, ShouldBeNil)
		So(
			string(data),
			ShouldEqual,
			string(expected),
		)
	})
}
