package main

import (
	"fmt"
	"time"
)

var c, python, java bool

func main() {
	// fmt.Printf("hello world %d\n", 5 + 1)
	// a, b := fun(2, 3)

	// fmt.Println(a, b)

	fmt.Println(time.Now())

	var i int
	fmt.Println(i, c, python, java)

	sum := 1
	/*for i := 0; i < 10; i++ {
		sum += i
	}*/
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func fun(x, y int) (int, int) {
	return x + y, x - y
}
