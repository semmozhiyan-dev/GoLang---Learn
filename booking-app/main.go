package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remaining Tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println("Welcome to ", conferenceName, " booking application")
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// its based on print formatting data we should use %v for variable
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("we have total of %v tickets and % v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//array
	bookings := []string{}

	//datatypes
	var userName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	//using pointer we get users input
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	//used for user input

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - (userTickets)
	//bookings[0] = userName + " " + lastName

	// slice
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The user value: %v\n", bookings)
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("slice length: %v\n", len(bookings))

	fmt.Println(remainingTickets)
	//pointer
	fmt.Println(&remainingTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v \n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, userTickets)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
