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
		fmt.Scan(&number)
		if number == 99 {
			fmt.Println(array)
			break
		}
		array = append(array, number)
	}
}
