package main

import (
	"fmt"
)

func main() {
	const x = 5
	const y = 5
	fmt.Println(x, "+", y, "=", Sum(x, y))
}

/*Sum is the sum function exported*/
func Sum(x int, y int) int {
	return x + y
}
