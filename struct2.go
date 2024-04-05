package main

import "cmp"

type CustomType interface {
	cmp.Ordered | string
}

type Foo[T CustomType] struct {
	Name string
	Age  int
	Data T
}
