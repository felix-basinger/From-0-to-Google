package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Gallows() string {
	rand.Seed(time.Now().UnixNano())
	var letterOfPlayer string

	var choiceOfWord int
	var gameWord string
	var words = [...]string{"корыто", "пластик", "енот", "сталактит", "гантеля"}

	choiceOfWord = rand.Intn(5)
	gameWord = words[choiceOfWord]
	var wordOfPlayer []string = make([]string, len(gameWord))
	count := 1
	countLetters := 0
	for count <= 5 {
		arrayGameWord := strings.Split(gameWord, "")
		fmt.Println("Enter your letter: ")
		fmt.Scan(&letterOfPlayer)

		if letterOfPlayer == arrayGameWord[countLetters] {
			wordOfPlayer[countLetters] = letterOfPlayer
			fmt.Println("The letter", countLetters+1, "is right")
			countLetters++
			if countLetters == len(arrayGameWord) {
				return "You win!"
			}
		} else {
			fmt.Println("The letter", countLetters+1, "is wrong( Try again")
			count++
		}
	}
	return "You lose("
}
