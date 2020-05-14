package main

import "fmt"

func addnumber(a int, b int) int {
	return a + b
}

func multiplereturn(a int, b int) (int, int) {
	return a, b
}

func main() {

	fmt.Println(addnumber(5, 10))
	//go does not support function overloading

	//go supports pass by value

	a, _ := multiplereturn(5, 10)
	fmt.Println(a)

	_, b := multiplereturn(5, 10)
	fmt.Println(b)

	c, d := multiplereturn(5, 10)
	fmt.Println(c, d)

	localfunc := addnumber
	fmt.Println(localfunc(5, 10))

	//adding func inside another function
	nestedfunc := func(h int, j int) int {
		return h + j
	}

	fmt.Println(nestedfunc(5, 10))

	//we can also pass a function to another function and return them
}
