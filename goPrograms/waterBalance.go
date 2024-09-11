package goPrograms

import (
	"fmt"
)

const waterOnKg float64 = 40

func WaterBalance() string {
	var weight float64
	var amountOfWater float64
	fmt.Println("Enter your weight: ")
	fmt.Scan(&weight)
	fmt.Println("How many liters of water you drink per day?: ")
	fmt.Scan(&amountOfWater)
	var waterNorma float64 = (weight * waterOnKg) / 1000
	if amountOfWater >= waterNorma-0.5 && amountOfWater <= waterNorma+0.5 {
		return "You drink enough"
	} else if amountOfWater < waterNorma-0.5 {
		fmt.Printf("You need to drink at least %.2f liters of water", waterNorma-0.5)
		return "You drink not enough"
	} else {
		fmt.Printf("You need to drink no more than %.2f liters of water", waterNorma+0.5)
		return "You drink too much water"
	}
}
