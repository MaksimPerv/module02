package main

import "fmt"

func main() {
	var A *int
	var B int = 8
	A = &B
	fmt.Println(*A)
	B = 10
	fmt.Println(*A)
}
