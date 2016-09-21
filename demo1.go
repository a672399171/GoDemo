package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("hello world %d", 5+1)
	a, b := fun(2, 3)

	fmt.Print(a)
	fmt.Print(b)

	fmt.Print(time.Now())
}

func fun(x, y int) (int, int) {
	return x + y, x - y
}
