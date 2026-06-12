package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",
		conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	for true {

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
		//user validation
		isvalidName := len(userName) >= 2 && len(lastName) >= 2 && strings.Contains(email, "@") && userTickets > 0 && userTickets <= remainingTickets
		// Ticket Validation
		if isvalidName {

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

			// Display bookings
			fmt.Printf("\nAll Bookings: %v\n", bookings)
			fmt.Printf("Bookings Type: %T\n", bookings)
			fmt.Printf("Number of Bookings: %v\n", len(bookings))

			// Extract first names
			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("First names of attendees: %v\n", firstNames)

			// Pointer Example
			fmt.Printf("Address of remainingTickets: %p\n", &remainingTickets)

			// Conference Sold Out
			if remainingTickets == 0 {
				fmt.Println("\nOur conference is fully booked!")
				fmt.Println("Come back next year.")
				break
			}

		} else if userTickets == remainingTickets {

			remainingTickets -= userTickets

		} else {

			fmt.Printf(
				"\nInvalid input. Please try again.\n",
			)

		}
	}
}
