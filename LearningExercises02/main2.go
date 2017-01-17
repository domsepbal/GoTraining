package main

import (
	"fmt"
	"strings"
)

func main() {
	//Divide a number in half and checks if it is even
	num := 10
	fmt.Println(divideandeven(num))

	//Shows how a slice will grow, once it hits cap, creates a new array and copies old on into it
	number := make([]int, 0, 5)
	for i := 0; i <= 10; i++ {
		number = append(number, i)
		fmt.Println(len(number), cap(number), i)
	}
}

func divideandeven (n int) (int, bool) {
	return n / 2, n % 2 == 0
}

func greatestnumber(n ...int) (int) {
	return 0
}
