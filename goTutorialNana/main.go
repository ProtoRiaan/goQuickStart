package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a totla of %v tickets and we still have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var userName string
		var firstName string
		var lastName string
		var email string
		var userTickets uint
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

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The following users have bookings: %v\nWe have %v tickets remaining\n\n", firstNames, remainingTickets)

		// fmt.Printf("We have %v tickets remaining\n", remainingTickets)

		// fmt.Printf("The whole Array : %v\n", bookings)
		// fmt.Printf("The first value of the Array: %v\n", bookings[0])

		// fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)

	}
}
