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

func reversCount(count int) {
	for i := count; i > 0; i-- {
		fmt.Println(i)
	}
}

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 10 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	evenNumber(50)
	reversCount(50)
	for {
		fmt.Println("Infinite loop")
		break // breaking the loop after one iteration
	}
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

}
