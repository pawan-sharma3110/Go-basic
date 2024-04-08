package main

import "fmt"

func evenNumber(number int) {
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println("Even Number")
		} else {
			fmt.Println(i)
		}
	}

}

// func reversCount(count int) {
// 		for i := count; i > 0; i-- {
// 		fmt.Println(i)
// 	}
// }
func reversCount(count int) {
    for i;i= count; i > 0; i-- {
        fmt.Println(i)
    }
}

func main() {
	// evenNumber(50)
	reversCount(50)
}
