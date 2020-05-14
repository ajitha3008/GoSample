package main

//we can have user defined errors as well

import (
	"errors"
	"fmt"
)

func divide(a int, b int) error {
	if b == 0 {
		return errors.New("cannot do, poda")
	}
	fmt.Println(a / b)
	return nil
}
func main() {
	a := 5
	b := 0
	err := divide(a, b)
	if err != nil {
		fmt.Println(err)
	}
}
