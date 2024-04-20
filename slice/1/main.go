package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	arr=append(arr, 20)
	fmt.Println(arr)
}
