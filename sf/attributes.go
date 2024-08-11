package sf

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrReflectionError = errors.New("sf: reflection error")
)

func AttributeAssign(data any, path string, value any) {
	attr := AttributeGet(data, path)
	attr.Value.Set(reflect.ValueOf(value))
}

type Attribute struct {
	Value reflect.Value
}

func AttributeGet(data any, path string) Attribute {
	// fmt.Println("Entering as", path)
	dot := strings.IndexRune(path, '.')
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		panic(ErrReflectionError)
	}
	strukt := rv.Elem()
	if strukt.Kind() != reflect.Struct {
		panic("not a struct")
	}

	if dot == -1 { // no dot found
		val := strukt.FieldByName(path)
		return Attribute{val}
	} else { // nested path
		field := strukt.FieldByName(path[:dot])
		nestedStrukt := field.Addr().Interface()
		return AttributeGet(nestedStrukt, path[dot+1:])
	}
}
