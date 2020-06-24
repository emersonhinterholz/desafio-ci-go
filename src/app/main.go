package main

import (
	"fmt"

	"../utils"
)

func main() {
	const x = 5
	const y = 5
	fmt.Println(x, "+", y, "=", utils.Sum(x, y))
}
