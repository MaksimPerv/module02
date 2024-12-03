package main

import "fmt"

func main() {
	weekend := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	work := make([]string, 5, 5)
	copy(work, weekend[0:5])
	fmt.Println(work)
	weekend = weekend[5:7]
	fmt.Println(weekend)
	dayOfTheWeek := append(work, weekend...)
	fmt.Println(dayOfTheWeek)
}
