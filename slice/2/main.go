package main

import "fmt"

var slice1 = []int{1, 2, 3, 4, 5}
var slice2 = []int{6, 7, 8, 9, 10}

func main() {
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}
