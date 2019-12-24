package main

import (
	"fmt"
)

const (
	first  = 1
	second = "second"
)

func main() {
	fmt.Println("Hello World")

	var i int
	i = 4
	fmt.Println(i)

	var p float32 = 2.2
	fmt.Println(p)

	firstName := "Mike"
	fmt.Println(firstName)

	l, m := complex(1, 2), complex(2, 1)
	fmt.Println(l, m)

	var secondName *string = new(string)
	*secondName = "Kevin"
	fmt.Println(*secondName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	const pi = 3.1416

	fmt.Println(3 + pi)
	fmt.Println(3.2 + pi)

	const c int = 3

	fmt.Println(3 + c)
	fmt.Println(3.2 + float32(c))

	fmt.Println(first, second)
}
