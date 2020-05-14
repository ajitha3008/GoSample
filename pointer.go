package main

import "fmt"

func main() {
	//var b *int
	//fmt.Println(*b) null  pointer deference error

	var c = new(int)
	fmt.Println(*c)
}
