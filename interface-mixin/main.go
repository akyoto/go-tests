package main

import (
	"fmt"
	"reflect"
)

type myInterface interface {
	Speak()
}

type embeddedType struct{}

func (t *embeddedType) Speak() {
	fmt.Println(reflect.TypeOf(t))
}

type rootType struct {
	embeddedType
}

func main() {
	var t myInterface
	t = &rootType{}
	t.Speak()
}
