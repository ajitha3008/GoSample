package main

import "fmt"

type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{1, "aj"}
	fmt.Println(g)
}

//also embedded struct and nested struct are possible
