package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
	fmt.Printf("%T", n)

	num := []float64{43, 56, 87, 12, 45, 57}
	ave := average(num...) //it will pass each item from the list one at a time over to the function
	fmt.Println(ave)

	fmt.Println(averageArray(num)) //Another way to write it for passing arguments into functions

	//Anon Func Expression

	greeting := func() {
		fmt.Println("Hello World!!")
	}
	greeting()

	//Callback functions
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

//Varadic function the ... means it can take infinite amount of arguments
func average(sf ...float64) (float64) {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}

func averageArray(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))

}
