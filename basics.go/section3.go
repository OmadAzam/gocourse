package main

import "fmt"

func main() {

	i := 0

	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}

	sum := 0

	for {
		sum = sum + 5
		fmt.Println(sum)
		if sum >= 55 {
			break
		}
	}

}
