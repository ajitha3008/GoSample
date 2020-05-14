package main

import "fmt"

func main() {
	in := make(chan int 1)//buffered channel
	out := make(chan int, 1)

	out <- 2

	select {
	case in <- 1:
		fmt.Println("wrote to in")
	case v := <-out:
		fmt.Println("wrote to out", v)
	}
}
