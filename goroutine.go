package main

//uses waitgroup from sync package to wait and finish
import (
	"fmt"
	"sync"
)

func printhello() {
	fmt.Println("Hello")
}
func main() {
	//go printhello()
	//time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		printhello()
		wg.Done()
	}()
	wg.Wait()
}
