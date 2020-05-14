package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	// no need of break in switch statements
	switch word {
	case "hi":
		fmt.Println("Hello")
	case "bi":
		fmt.Println("Bubye")
		fallthrough
	case "any", "bunny":
		fmt.Println("Any other statement")
	default:
		fmt.Println("I am default")
	}
}
