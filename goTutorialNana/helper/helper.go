package helper

import (
	"fmt"
	"strconv"
	"strings"
)

var Bookings = make([]map[string]string, 0)

// create a map or a user

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTickets

}

func BookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) uint {

	var userData = make(map[string]string)

	type UserData struct {
		firstName       string
		lastName        string
		email           string
		numberOfTickets uint
		isVerified      bool
	}

	remainingTickets = remainingTickets - userTickets
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticketNumber"] = strconv.FormatUint(uint64(userTickets), 10)
	Bookings = append(Bookings, userData)

	fmt.Printf("Than you %v %v for booking %v tickets. You will receive confirmation at %v\n", firstName, lastName, userTickets, email)

	// firstNames := []string{}
	// for _, booking := range Bookings {
	// 	var names = strings.Fields(booking)
	// 	firstNames = append(firstNames, names[0])
	// }

	firstNames := []string{}
	for _, booking := range Bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	fmt.Printf("The following users have bookings: %v\nWe have %v tickets remaining\n\n", firstNames, remainingTickets)
	return remainingTickets
}
