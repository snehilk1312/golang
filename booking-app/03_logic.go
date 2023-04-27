package main

import "fmt"

func main() {

	var conferenceName string = "go conference" // variable name must be used
	const conferenceTickets uint = 50           // constant
	var remainingTickets int = 50

	fmt.Printf("Welcome to our conferece %v booking app\n", conferenceName)
	fmt.Printf("Total No of Tickets: %v\n", conferenceTickets)

	// user input

	var firstName string
	var lastName string
	var email string
	var numTcks int
	// var bookings [50]string // defining an array var
	var bookings []string // defining an slice var

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter num of tickets you want: ")
	fmt.Scan(&numTcks)

	fmt.Printf("congratulations , %s %s you have %v tickets.Please check %s for more detail\n", firstName, lastName, numTcks, email)

	remainingTickets = remainingTickets - numTcks

	// bookings[0] = firstName + " " + lastName //when array
	// bookings[4] = firstName + " " + lastName //when array

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole booking array: %v\n", bookings)
	fmt.Printf("The first element of booking array: %v\n", bookings[0])
	fmt.Printf("The type of  booking array: %T\n", bookings)
	fmt.Printf("Length booking array: %v\n", len(bookings))

	fmt.Printf("Remaining No of Tickets: %v\n", remainingTickets)

}
