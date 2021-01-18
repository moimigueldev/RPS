package main

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = [3]string{"rock", "paper", "scissors"}

type stats struct {
	wins, losts, draws int
}

var (
	user string
	cmp  string
)

var gameStats = stats{}

func main() {
	getUserChoice()
	setCMPChoice()
	compare()

}

func setCMPChoice() string {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	cmp = choices[y1.Intn(3)]

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
		fmt.Println("You Lost")
		gameStats.losts++
	} else if user == "paper" && cmp == "rock" {
		fmt.Println("You Won")
		gameStats.wins++
	} else if user == "scissors" && cmp == "paper" {
		fmt.Println("You Won")
		gameStats.wins++
	} else if user == "paper" && cmp == "scissors" {
		fmt.Println("You Lost")
		gameStats.losts++
	} else if user == "rock" && cmp == "scissors" {
		fmt.Println("You Won")
		gameStats.wins++
	} else if user == "scissors" && cmp == "rock" {
		fmt.Println("You Lost")
		gameStats.losts++
	} else {
		fmt.Println("You tied")
		gameStats.draws++
	}

}
