package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	//-----------------------------------------VARIABLES----------------------------------------
	var i int = 10
	fmt.Println(i)

	var j = 11
	fmt.Println(j)

	k := 12 //declaration format can be used only in function
	fmt.Println(k)

	var l int //undefined numerical variable defaults to 0
	fmt.Println(l)

	l = 999
	fmt.Println(l)

	//unused variables in Go will throw an error

	//typecasting
	var a int64 = 5
	var b float32 = 5.5
	fmt.Println(float32(a) + b)
	fmt.Println(a + int64(b))

	var s string
	s = "Ajitha" //interpreted literal
	fmt.Println(s)

	s = `ajsjd
	sdsdf
	dsf`
	fmt.Println(s) //raw string literal

	mystring := "Hello"
	g := mystring[0]
	h := mystring[4]
	sub := mystring[0:3]
	sub1 := mystring[:3]
	fmt.Println(mystring, g, h, string(g), string(h), sub, sub1)

	fmt.Println("Length=", len(mystring))

	var r rune = 127757 // unicode datatype
	fmt.Println(string(r))

	//arrays
	var vals [3]int
	vals[0] = 1
	vals[1] = 2
	vals[2] = 3
	//vals[4]=5 throws error

	fmt.Println(vals, vals[1])
	//-----------------------------------------CONDITIONAL STATEMENTS----------------------------------------
	z := 10
	// 0 is nt equal to false
	if z > 5 {
		fmt.Println("z is bigger than 5")
	} else {
		fmt.Println("z is lesser than 5")
	}

	for l := 0; l < 7; l++ {
		if l == 3 {
			break
		}
		fmt.Println(l)
	}

	m := "Hello"
	for k, v := range m {
		fmt.Println(k, v, string(v))
	}

}
