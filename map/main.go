package main

import "fmt"

func main() {
	a := map[string]map[string][]string{}
	a["Иванов"] = map[string][]string{
		"Золотая рыбка":      {"Первое", "Второе"},
		"Мастер и Маргарита": {"Третье"},
	}
	for _, v := range a {
		fmt.Println(len(v))
		fmt.Println(v)
	}
}
