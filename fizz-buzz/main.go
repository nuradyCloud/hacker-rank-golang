package main

import (
	"fmt"
)

func main() {

	var number int32
	i := int32(1)

	fmt.Print("Input number integer: ")
	fmt.Scan(&number)
	fmt.Print("Result: \n")

	for i = 1; i <= number; i++ {
		fizzBuzz(i)
	}

}

func fizzBuzz(n int32) {

	fizz := "Fizz"
	buzz := "Buzz"

	if n%3 == 0 && n%5 == 0 {
		fmt.Println(fizz + buzz)
	} else if n%3 == 0 {
		fmt.Println(fizz)
	} else if n%5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(n)
	}

}
