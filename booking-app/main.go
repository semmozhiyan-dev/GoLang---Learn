package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	//FUNC CALL
	greetUsers()

	for true {
		//FUNC CALL
		userName, lastName, email, userTickets := getUserInput()
		//user validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userName, lastName, email, userTickets) // Ticket Validation
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userName, lastName, email, userTickets)

			// Extract first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of our bookings are: %v\n", firstNames)

			// Pointer Example
			fmt.Printf("Address of remainingTickets: %p\n", &remainingTickets)

			// Conference Sold Out
			if remainingTickets == 0 {
				fmt.Println("\nOur conference is fully booked!")
				fmt.Println("Come back next year.")
				break
			}

		} else {
			if !isValidName || !isValidEmail || !isValidTicketNumber {
				fmt.Printf(
					"\nInvalid name or email. Please try again.\n",
				)
			}

			fmt.Printf(
				"\nInvalid input. Please try again.\n",
			)

		}
	}
	city := "london"
	switch city {
	case "new york":
		fmt.Println("Welcome to New York!")
	case "singapore":
		fmt.Println("Welcome to Singapore!")
	case "london":
		fmt.Println("Welcome to London!")
	default:
		fmt.Println("Welcome to our conference!")
	}

}
func greetUsers() {

	fmt.Printf("Welcome to %v our conference booking application!\n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("First names of bookings: %v\n", firstNames)
	return firstNames
}

func validateUserInput(userName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var userName string
	var lastName string
	var email string
	var userTickets uint

	// User Input
	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return userName, lastName, email, userTickets
}

func bookTicket(userName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	// Add booking to slice
	bookings = append(bookings, userName+" "+lastName)

	// Booking Confirmation
	fmt.Printf(
		"\nThank you %v %v for booking %v tickets.\n",
		userName,
		lastName,
		userTickets,
	)

	fmt.Printf(
		"You will receive a confirmation email at %v.\n",
		email,
	)

	fmt.Printf(
		"%v tickets remaining for %v.\n",
		remainingTickets,
		conferenceName,
	)

}
