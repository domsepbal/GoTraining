package main

import "fmt"

func main() {
	for i := 1; i <= 200; i++{
		fmt.Printf("%v - %v - %v \t", i, string(i), []byte(string(i)))
		if(i % 3 == 0){
			fmt.Println()
		}
	}
}
