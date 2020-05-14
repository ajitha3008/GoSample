package main

import "fmt"

func main() {

	in := make(chan string)
	out := make(chan string)

	go func() {
		message := <-in
		out <- fmt.Sprintf("Message is :" + message)
	}()
	in <- "Hello"
	close(in)
	name := <-out
	fmt.Println(name)
}

//buffered channel is possible
//select helps us to work with multiple goRoutines
