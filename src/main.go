package main

import "fmt"

type AmericanVelocity float64

type EuropeanVelocity float64

func main() {
	// Константы для преобразования
	const mpsToKmph = 3.6    // метры в секунду в километры в час
	const mpsToMph = 2.23694 // метры в секунду в мили в час

	// Преобразование скорости 120.4 м/сек в км/ч
	var speedInMps1 float64 = 120.4
	var speedInKmph EuropeanVelocity = EuropeanVelocity(speedInMps1 * mpsToKmph)

	// Преобразование скорости 130 м/с в миль/ч
	var speedInMps2 float64 = 130
	var speedInMph AmericanVelocity = AmericanVelocity(speedInMps2 * mpsToMph)

	// Вывод результатов
	fmt.Printf("Скорость %.2f м/с равна %.2f км/ч\n", speedInMps1, speedInKmph)
	fmt.Printf("Скорость %.2f м/с равна %.2f миль/ч\n", speedInMps2, speedInMph)
}
