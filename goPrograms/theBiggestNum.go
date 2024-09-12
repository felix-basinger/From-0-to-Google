package goPrograms

import (
	"fmt"
	"math/rand"
	"time"
)

func RandArray() []int {
	rand.Seed(time.Now().UnixNano())

	var quantity int
	var myArray []int

	fmt.Println("Enter the quantity of numbers in array: ")
	fmt.Scan(&quantity)

	myArray = make([]int, quantity)

	for i := 0; i < len(myArray); i++ {
		myArray[i] = rand.Intn(1000)
	}
	return myArray
}

func TheBiggestNum(giveArray func() []int) int {
	array := giveArray()
	var maxNum = array[0]
	var index int
	for i := 1; i < len(array); i++ {
		if array[i] > maxNum {
			maxNum = array[i]
			index = i
		}
	}
	fmt.Printf("The biggest number in array is %d, index %d", maxNum, index)
	return maxNum
}
