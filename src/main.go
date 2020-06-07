package main

import (
	"fmt"
)

func main() {
	const x = 5
	const y = 5
	fmt.Println(x, "+", y, "=", sum(x, y))
}

func sum(x int, y int) int {
	return x + y
}
