package main

import "fmt"

func main() {
	workingWithSlices()

}

func workingWithSlices() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 10, 11, 12)

	fmt.Println(slice2)

	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2]
	fmt.Println(slice3, slice4, slice5)
}
