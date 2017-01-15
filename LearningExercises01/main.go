package main

import "fmt"

func main() {

	//Exercise to print multiples of 3 and replace with fizz, multiples of 5 with buzz, multiples of both with fizzbuzz
	for i := 1; i <= 100 ; i++  {

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(i, " --FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println(i, " --Fizz")
		} else if i % 5 == 0 {
			fmt.Println(i, " --Buzz")
		} else {
			fmt.Println(i)
		}
	}
	//Separation
	fmt.Println()

	//Find the sum of all multipes of 3 or 5
	multipleOfThree := 0
	multipleOfFive := 0
	counter := 0

	for i := 1; i < 1000; i++ {
		if i % 3 == 0 {
			multipleOfThree += i
			counter += i
		} else if i % 5 == 0 {
			multipleOfFive += i
			counter += i
		}
	}

	fmt.Printf("Sum of multiples of three: %v\n", multipleOfThree)
	fmt.Printf("Sum of multiples of five: %v\n", multipleOfFive)
	fmt.Printf("Sum of multiples of both: %v\n", counter)

	//Print all even numbers between 0 and 100
	for i := 0 ; i <=100 ; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	//Separation
	fmt.Println()

	//Get a name and input into a greeting
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("\nHello, %v!", name)
}
