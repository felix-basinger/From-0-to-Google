package games

import (
	"fmt"
	"math/rand"
	"time"
)

func ScisRockPaper() string {
	rand.Seed(time.Now().UnixNano())

	var rounds int
	var countComp int
	var countPlayer int
	countOfRounds := 1
	var trial int
	var choiceOfPlayer int
	var choiceOfComp int

	items := [3]string{"Paper", "Rock", "Scissors"}

	fmt.Println("Welcome to the rock paper!")
	fmt.Println("How many rounds do you wanna play?")
	fmt.Scan(&rounds)

	for trial <= rounds-1 {
		fmt.Println("\nEnter your item: ")
		fmt.Printf("1 - %s\n2 - %s\n3 - %s\n", items[0], items[1], items[2])
		fmt.Scan(&choiceOfPlayer)
		for choiceOfPlayer > 3 || choiceOfPlayer < 1 {
			fmt.Println("You neeed to choose 1, 2 or 3 ")
			fmt.Printf("1 - %s\n2 - %s\n3 - %s\n", items[0], items[1], items[2])
			fmt.Scan(&choiceOfPlayer)
		}

		choiceOfComp = rand.Intn(3)

		switch choiceOfPlayer {
		case 1:
			if choiceOfComp == 2 {
				fmt.Println("\nRound ", countOfRounds, " win!")
				fmt.Printf("%s -> %s\n", items[0], items[1])
				countPlayer++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else if choiceOfComp == 1 {
				fmt.Println("\nRound ", countOfRounds, " draw!")
				fmt.Printf("%s -> %s\n", items[0], items[0])
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else {
				fmt.Println("\nRound ", countOfRounds, " lose!")
				fmt.Printf("%s -> %s\n", items[2], items[1])
				countComp++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			}
		case 2:
			if choiceOfComp == 3 {
				fmt.Println("\nRound ", countOfRounds, " win!")
				fmt.Printf("%s -> %s\n", items[1], items[2])
				countPlayer++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else if choiceOfComp == 2 {
				fmt.Println("\nRound ", countOfRounds, " draw!")
				fmt.Printf("%s -> %s\n", items[1], items[1])
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else {
				fmt.Println("\nRound ", countOfRounds, " lose!")
				fmt.Printf("%s -> %s\n", items[0], items[1])
				countComp++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			}
		case 3:
			if choiceOfComp == 1 {
				fmt.Println("\nRound ", countOfRounds, " win!")
				fmt.Printf("%s -> %s\n", items[2], items[0])
				countPlayer++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else if choiceOfComp == 3 {
				fmt.Println("\nRound ", countOfRounds, " draw!")
				fmt.Printf("%s -> %s\n", items[2], items[2])
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			} else {
				fmt.Println("\nRound ", countOfRounds, " lose!")
				fmt.Printf("%s -> %s\n", items[1], items[2])
				countComp++
				fmt.Printf("Computer %d - %d Player\n", countComp, countPlayer)
			}

		}
		countOfRounds++
		trial++
	}
	if countPlayer < countComp {
		return "\nComputer win!"
	} else if countPlayer == countComp {
		return "\nDraw!"
	} else {
		return "\nYou win!"
	}
}
