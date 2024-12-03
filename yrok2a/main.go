package main

import (
	"fmt"
	"math"
)

func main() {
	// Длина окружности
	circumference := 35.0

	// Находим радиус окружности по формуле R = C / (2 * π)
	radius := circumference / (2 * math.Pi)

	// Объявляем указатель на радиус
	R := &radius

	// Вычисляем площадь круга по формуле S = π * R^2
	area := math.Pi * math.Pow(*R, 2)

	// Округляем площадь до двух знаков после запятой
	area = math.Round(area*100) / 100

	// Выводим результат
	fmt.Printf("Радиус окружности: %.2f\n", *R)
	fmt.Printf("Площадь круга: %.2f\n", area)
}
