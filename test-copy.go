package main

import (
	"fmt"
)

func main() {
	a := []string{"1", "2"}
	b := []string{"3", "4"}
	copy(b, a)
	fmt.Println(b)
}
