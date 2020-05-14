package main

import (
	"fmt"
)

//slice example

func main() {
	fmt.Println("Hello, playground")
	s := make([]string, 0)
	fmt.Println(len(s))
	s = append(s, "hello") //adds new element at the end
	fmt.Println("after appending", s)
	s[0] = "world"
	fmt.Println("after changing", s)
	fmt.Println(s[0])
	fmt.Println(len(s))
	//fmt.Println(s[1]) throws error

	mymap := make(map[string]int)
	mymap["hello"] = 300

}
