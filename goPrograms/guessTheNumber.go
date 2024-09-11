package goPrograms

import (
	"fmt"
	"math"
)

func ArrayGenerator() []int {
	var countOfNumbers int
	fmt.Println("Enter quantity of numbers that you want to guess for me: ")
	fmt.Scan(&countOfNumbers)              // Scan - это запись ввода пользователя из консоли в переменную
	myArray := make([]int, countOfNumbers) // Для генерации массива из N чисел используем срезы
	for i := 0; i < countOfNumbers; i++ {
		myArray[i] = i + 1
	}
	return myArray
}

func GuessMyNumber(giveArray func() []int) string { // В Go функция может передаваться в качестве параметра другой функции
	var attempt = 1                    // Переменная-счетчик
	var myArray = giveArray()          // Использование функции из параметра
	var low = 0                        // Нижняя граница поиска
	var high = myArray[len(myArray)-1] // Верхняя граница поиска. Выясняем длинну массива (количество элементов) с помощью len()
	var yesNo int
	var log2 = math.Ceil(math.Log2(float64(len(myArray)))) // Ceil - это функция наименьшее возвращает число, которое больше либо равно заданому значению
	// Log2 - это функция, выисляющая логарифм от N по основанию 2.
	// Ее мы используем для указания максимального количества попыток, за которые будет отгадано число
	fmt.Printf("\nWrite your number on a paper. I will try to guess it less than in %.0f attempts!\n", log2) // Передаем переменную количества попыток в строку
	// при помощи F-строки и округления до целого числа

	// Алгоритм бинарного поиска имеет сложность O(log n), где n - это количество элементов в отсортированном массиве
	// При помощи данной сложности мы и находим максимальное количество попыток, за которые программа угадывает число
	// Алгоритм бинарного поиска на примере >>>
	for low <= high {
		var mid = (low + high) / 2 // Если нижняя граница поиска меньше верхней, то целочисленно делим их сумму на 2, чтобы вполовину сократить радиус поиска
		var guess = myArray[mid]   // Записываем это значение в индекс и выводим результат - количество всех элементов / 2
		fmt.Printf("\nAttempt № %d\n%d is this your number?\n\n", attempt, guess)
		attempt++ // Увеличиваем счетчик при помощи инкримента
		fmt.Println("1 - Yes\n0 - No")
		_, err := fmt.Scan(&yesNo)                     // 1 уровень защиты от дураков. Отлавливаем ошибку ввода. Если введено не целое число, то фиксируем ошибку
		for err != nil || (yesNo != 1 && yesNo != 0) { // 2 уровень защиты от дураков. Если введено любое число или символ, кроме 1 или 0 -
			// отправляем в бесконечный цикл, пока не введет нужное значение
			fmt.Println("\nPrint 1 or 0: ")
			fmt.Println("1 - Yes\n0 - No")
			fmt.Scan(&yesNo)
		}
		// Проверяем, если число угадано и пользователь ввел 1, то выходим из цикла и завершаем программу
		// Если же пользователь ввел 0, то продолжаем поиск, уточняя меньше ли загаданное число названного числа (guess) или больше
		// Если код верный и алгоритм написан правильно, программа должна затрачивать не больше указанного количества попыток
		switch yesNo { // Условная конструкция switch/case - заменитель для тех, кому надоел if/else. А так по сути между ними разницы нет никакой
		case 1:
			if yesNo == 1 {
				fmt.Printf("\nYour number is %d\n\n", guess)
				return "I guessed!"
			}
		case 0:
			if yesNo == 0 {
				fmt.Printf("\nIs your number greater than %d?\n\n", guess)
				fmt.Println("1 - Yes\n0 - No")
				_, err := fmt.Scan(&yesNo)
				for err != nil || (yesNo != 1 && yesNo != 0) {
					fmt.Println("\nPrint 1 or 0: ")
					fmt.Println("1 - Yes\n0 - No")
					fmt.Scan(&yesNo)
				}
				if yesNo == 1 {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		}
	}
	return "I guessed!"
}
