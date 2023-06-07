package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number int32

	fmt.Print("Input number integer: ")
	fmt.Scan(&number)
	fmt.Printf("Binary representation of %d is: %s.\n", number, strconv.FormatInt(int64(number), 2))
	fmt.Printf("The total number of set bits in %d is %d.\n", number, countBits(number))

}

func countBits(n int32) int32 {

	var count int32
	count = 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count

}
