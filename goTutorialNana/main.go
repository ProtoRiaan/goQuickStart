package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a totla of %v tickets and we still have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
