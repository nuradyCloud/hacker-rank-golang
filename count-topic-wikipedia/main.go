package main

import (
	"fmt"
)

func main() {

	var topic string
	fmt.Print("Input topic search wikipedia: ")
	fmt.Scan(&topic)
	fmt.Print("Result count topic: \n")
	//result := compressedString(message)
	//fmt.Printf("%s", result)

}

/*
 * Complete the function below.
 * https://en.wikipedia.org/w/api.php?action=parse&section=0&prop=text&format=json&page=[topic]
 */
func getTopicCount(topic string) int {

	return 0
}
