package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	A string
	B int
	C C
}

type C struct {
	X string
}

func main(){
	c := C{
		X : "nihao",
	}

	t := Test{
		A : "你好",
		B : 1,
		C : c,
	}

	res := reflect.ValueOf(&t).Elem()

	fmt.Println("结果1", res.FieldByName("A"))

	typ := reflect.TypeOf(t)

	for i := 0; i < res.NumField(); i++ {
		piece := typ.Field(i)
		fmt.Println("结果xx", piece.Name)
	}
}