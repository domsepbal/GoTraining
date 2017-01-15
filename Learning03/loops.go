package main

import "fmt"

func main() {
	for i := 0; i <= 30; i++{
		for l := 30; l > i+2 ; l-- {
			fmt.Print("\t")
		}
		//Left side
		for k := 1; k < i - 1 ; k++  {
			fmt.Print("*")
		}
		//Right Side
		for j := 0; j < i - 1 ; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}

	for i := 30; i >= 0 ; i-- {
		for j := 0;j < i ; j++  {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

/*fmt.Printf("%v - %v - %v \t", i, string(i), []byte(string(i)))
if(i % 3 == 0){
fmt.Println()*/
