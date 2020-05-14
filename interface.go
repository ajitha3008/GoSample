package main

import "fmt"

func main() {
	var inter interface{}
	inter = "hello"
	i := inter.(string)
	k, ok := inter.(int)
	fmt.Println(inter, i, k, ok)
}
