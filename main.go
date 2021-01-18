package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

var choices = [3]string{"rock", "paper", "scissors"}

var play = true

type stats struct {
	wins, loses, draws int
}

var (
	user string
	cmp  string
)

var gameStats = stats{}

func main() {

	for play {

		getUserChoice()
		setCMPChoice()
		compare()
		printStats()
		playAgain()

	}
}

func setCMPChoice() string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	cmp = choices[y1.Intn(3)]

	fmt.Printf("\n\nComputer Chose: %v", cmp)

	return cmp
}

func getUserChoice() string {
	fmt.Printf("Make your choice: (1) Rock (1) Paper (3) Scissors : ")
	fmt.Scanln(&user)

	switch user {
	case "1":
		user = "rock"
	case "2":
		user = "paper"
	case "3":
		user = "scissors"
	default:
		fmt.Println("Please Choose a valid option")
		getUserChoice()
	}

	return user
}

func compare() {

	if user == "rock" && cmp == "paper" {
		fmt.Printf("\n\n~YOU LOST~\n\n")
		gameStats.loses++
	} else if user == "paper" && cmp == "rock" {
		fmt.Printf("\n\n~YOU WON~\n\n")
		gameStats.wins++
	} else if user == "scissors" && cmp == "paper" {
		fmt.Printf("\n\n~YOU WON~\n\n")
		gameStats.wins++
	} else if user == "paper" && cmp == "scissors" {
		fmt.Printf("\n\n~YOU LOST~\n\n")
		gameStats.loses++
	} else if user == "rock" && cmp == "scissors" {
		fmt.Printf("\n\n~You Won~\n\n")
		gameStats.wins++
	} else if user == "scissors" && cmp == "rock" {
		fmt.Printf("\n\n~YOU LOST~\n\n")
		gameStats.loses++
	} else {
		fmt.Printf("\n\n~YOU TIED~\n\n")
		gameStats.draws++
	}
}

func playAgain() {
	input := "tempt"
	fmt.Printf("Do You want To Play Again? (1) Yes (2) No > ")
	fmt.Scanln(&input)
	switch input {
	case "2":
		play = false
		clearTheScreen()
		printStats()
	default:
		play = true
		clearTheScreen()
	}

}

func clearTheScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	// Animate the time always in the same position
	fmt.Println(time.Now())
	i := 200
	time.Sleep(time.Duration(i) * time.Millisecond)
}

func printStats() {

	if play {
		fmt.Printf("\n| WINS: %v | DRAWS: %v  | lOSES: %v |\n\n", gameStats.wins, gameStats.draws, gameStats.loses)
	} else {
		fmt.Printf("\nTHANKS FOR PLAYING, HERE ARE YOUR STATS: ")
		fmt.Printf("\n\n| WINS: %v | DRAWS: %v  | lOSES: %v |\n\n", gameStats.wins, gameStats.draws, gameStats.loses)
	}
}
