// code with the syntax error
// fix the syntax error by adding a semicolon at the end

package main

import (
	"fmt"
)

func main() {
	array := []int{}
	var number int

	for {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&number)
		if number == 99 {
			fmt.Println("array is ", array)
			fmt.Println("length of array is ", len(array))
			break
		}
		array = append(array, number)
	}

}
