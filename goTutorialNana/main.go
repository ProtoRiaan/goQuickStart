package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var userName string
var firstName string
var lastName string
var email string
var userTickets uint

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a totla of %v tickets and we still have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(helper.Bookings) < 50 {
		// var userName string
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint
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

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		// booking logic
		// if userTickets > remainingTickets{
		// 	fmt.Println("We dont have that many tickets available, please try booking less tickets")
		// 	// break
		// 	continue
		// }
		if isValidEmail && isValidName && isValidTickets {

			remainingTickets = helper.BookTickets(remainingTickets, userTickets, firstName, lastName, email)
			go sendTicket()

			if !(remainingTickets > 0) {
				// end program
				fmt.Println("Our conference is booked out, better luck next time")
				break
			}

			// fmt.Printf("We have %v tickets remaining\n", remainingTickets)

			// fmt.Printf("The whole Array : %v\n", bookings)
			// fmt.Printf("The first value of the Array: %v\n", bookings[0])

			// fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
		} else {
			fmt.Println("Your input data is invalid")
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is invalid")
			}
			if !isValidTickets {
				fmt.Println("We are not able to process the number of tickets you requested")
			}
		}
	}
}

func sendTicket() {
	time.Sleep(10 * time.Second)
	var ticketemail = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticketemail, email)
	fmt.Println("#########################")

}
