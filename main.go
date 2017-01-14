package main
import (
	"fmt"
)


func main() {



	const (
	_ = iota //Throw away the zero
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	)


	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)






	/*	res, _ := http.Get("http://www.ffxiah.com/browse/1/hand-to-hand")
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%s", page)*/

	/*Global Declaration and another way to init variables
	var(
		f = 3
		g = 6
	)*/

	/*	a := 10 //Short hand method init
		var b int //0 value

		fmt.Println("Hello World")
		fmt.Printf("a: %v \n", a) //%v figures out what to print
		fmt.Printf("b: %v \n", b)
		fmt.Printf("b: %T \n", b) //%T prints out the type
		fmt.Printf("Global Variables: %v, %v \n", f, g)



		//Input from User
		fmt.Print("Enter a number to be doubled: ")
		var input float64
		fmt.Scanf("%f", &input)

		output := input * 2

		fmt.Printf("You number doubled: %v ", output)*/
}
