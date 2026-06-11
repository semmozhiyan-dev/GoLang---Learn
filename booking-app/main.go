package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remaining Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to ", conferenceName, " booking application")
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// its based on print formatting data we should use %v for variable
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("we have total of %v tickets and % v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//datatypes
	var userName string
	var userTickets int
	//ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("USER %v BOOKED %v TICKETS.\n", userName, userTickets)

}
