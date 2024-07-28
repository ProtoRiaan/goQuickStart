package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a totla of %v tickets and we still have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name

	fmt.Println("Enter your username:")
	fmt.Scan(&userName)
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Than you %v %v for booking %v tickets. You will receive confirmation at %v\n", firstName, lastName, userTickets, email)

	fmt.Println(&remainingTickets)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
	fmt.Printf("User %v booked %v tickets for %v\n", userName, userTickets, conferenceName)
}
