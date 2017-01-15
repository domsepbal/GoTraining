package main

import "fmt"

func main() {
	width := 20
	height := 40

	for i := 0; i <= height; i++{
		//Right Side
		for j := width; j < i - 1 ; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

	for i := height; i >= 0 ; i-- {
		for j := width;j < i ; j++  {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

/*fmt.Printf("%v - %v - %v \t", i, string(i), []byte(string(i)))
if(i % 3 == 0){
fmt.Println()*/
