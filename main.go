package main

import (
	"fmt"

	_ "github.com/felixge/fgprof"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	println("Hello Lina")
}
