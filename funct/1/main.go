package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	fmt.Println("Your serquence")
	for i := 0; i < n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}

}
func main() {
	var n int
	fmt.Println("enter hte serqunce last position :")
	fmt.Scan(&n)
	fibonacci(n)
}
