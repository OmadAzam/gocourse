package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	numbers := []int{2, 4, 6, 8, 10}

	for _, n := range numbers {
		fmt.Println(n)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd number:", i)
		if i == 5 {
			break
		}

	}

	rows := 5
	for i := 1; i <= 5; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := range 10 {
		fmt.Println(10 - i)
	}
}
