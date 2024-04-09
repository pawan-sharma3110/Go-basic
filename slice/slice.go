package main

import (
	"fmt"
)

var list []string

var names = make([]string, 10)
var try = []int{1, 2, 3, 4}

func main() {
	names[0] = "Pawan"
	names[1] = "vikash"
	names[2] = "arun"
	names[3] = "ashu"
	names[4] = "pardeep"
	names[5] = "usha"
	names[7] = "ashu"
	list = append(list, "kuch")
	fmt.Println(list)

	list = append(list, "ashu001", "001kuc00111h")
	fmt.Println(list)

	for i, v := range list {
		fmt.Println(i, v)
	}

	for i, v := range names {
		fmt.Println(i, v)
	}
	fmt.Println(try)

}
