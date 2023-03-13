package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var game bool = true
	var choice string
	fmt.Printf("Welcome to Rock-paper-scissors")
	var score int = 0

	for game {
		fmt.Print("\nScore:", score, "\n1) Rock\n2) Paper\n3) Scissors\n4) Exit\nChoose any\n>  ")
		fmt.Scanln(&choice)
		var botChoice int = rand.Intn(3)
		if choice == "1" && botChoice == 1 { // draw
			fmt.Println("You : Rock\nRival : Rock\nDraw! Next round. ")
		} else if choice == "2" && botChoice == 2 {
			fmt.Println("You : Paper\nRival : Paper\nDraw! Next round. ")
		} else if choice == "3" && botChoice == 3 {
			fmt.Println("You : Scissors\nRival : ScissorsnDraw! Next round. ")
		} else if choice == "1" && botChoice == 3 { // win
			score += 1
			fmt.Println("You : Rock\nRival : Scissors\nYou won! Next round. ")
		} else if choice == "2" && botChoice == 1 {
			score += 1
			fmt.Println("You : Paper\nRival : Rock\nYou won! Next round. ")
		} else if choice == "3" && botChoice == 2 {
			score += 1
			fmt.Println("You : Scissors\nRival : Paper\nYou won! Next round. ")
		} else if choice == "1" && botChoice == 2 { // lose
			score -= 1
			fmt.Println("You : Rock\nRival : Paper\nYou lose! Next round. ")
		} else if choice == "2" && botChoice == 3 {
			score -= 1
			fmt.Println("You : Paper\nRival : Scissors\nYou lose! Next round. ")
		} else if choice == "3" && botChoice == 1 {
			score -= 1
			fmt.Println("You : Scissors\nRival : Rock\nYou lose! Next round. ")
		} else if choice == "4" {
			main()
		}

		if choice > "4" {
			fmt.Println("Invalid input!")
		}
	}
}
