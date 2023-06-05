package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var message string

	fmt.Print("Input string: ")
	fmt.Scan(&message)
	fmt.Print("Result: \n")
	result := compressedString(message)
	fmt.Printf("%s", result)
}

func compressedString(message string) string {

	compressedString := ""
	messageArray := strings.Split(message, "")
	strLen := len(message)
	for i := 0; i < strLen; i++ {
		count := 1
		currentLetter := messageArray[i]
		for i < strLen-1 && messageArray[i] == messageArray[i+1] {
			count++
			i++
		}
		if count == 1 {
			compressedString += currentLetter
		} else {
			compressedString += currentLetter + strconv.Itoa(int(count))
		}
	}
	return compressedString
}
