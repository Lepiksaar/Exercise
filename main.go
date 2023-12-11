package main

import (
	"fmt"

	"github.com/Lepiksaar/Exercise/first"
)

func main() {
	array := []int{1, 4, 3, 6, 89, 45, 234, 57, 769, 99}
	fmt.Print(first.Mergesort(array))
}
