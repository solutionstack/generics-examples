package main

import "fmt"

func Print[T any](arg []T) []T {

	for _, v := range arg {
		fmt.Println(v)
	}

	return arg
}

func main() {

	a := []string{"a", "b", "c", "de"}
	b := []int{1, 2, 3, 4}
	Print(a)
	Print(b)

	//generic list
	list := LinkedList[string]{}
	list.Insert("aa")
	list.Insert("bb")
	list.Insert("cc")
	list.Insert("dd")
	list.Print()

	//cust struct

	c := Foo[string]{
		Name: "a",
		Age:  0,
		Data: "ss",
	}
	fmt.Printf("%v", c)
}
