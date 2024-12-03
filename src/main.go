package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "104"
	b := 35
	a1, _ := strconv.Atoi(a)
	b1 := strconv.Itoa(b)
	fmt.Println(a1, b1)
}
