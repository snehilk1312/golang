package main

import "fmt"

func main() {

	var conferenceName string = "go conference" // variable name must be used
	const conferenceTickets uint = 50           // constant
	var remainingTickets int = 50

	fmt.Printf("Welcome to our conferece %v booking app\n", conferenceName)
	fmt.Printf("Total No of Tickets: %v\nAvailable No of Tickets: %v\n", conferenceTickets, remainingTickets)

	// printing data type
	fmt.Printf("data type of conferenceTickets: %T\ndata type of remainingTickets: %T\ndata type of conferenceName: %T\n", conferenceTickets, remainingTickets, conferenceName)

	// user input

	var firstName string
	var lastName string
	var email string
	var numTcks int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter num of tickets you want: ")
	fmt.Scan(&numTcks)

	fmt.Printf("congratulations , %s %s you have %v tickets.Please check %s for more detail\n", firstName, lastName, numTcks, email)

}
