package main

import (
	"fmt"
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
}
