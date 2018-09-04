package dao

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"testing"
)

func TestToCamel(t *testing.T) {
	fmt.Println(strcase.ToCamel("book_id"))
}
