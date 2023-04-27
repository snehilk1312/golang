package main

import "fmt"

func main() {
	//fmt.Print("Hello World!")   // prints without newline
	fmt.Println("Hello World!")          // prints with newline
	var conferenceName = "go conference" // variable name must be used
	const conferenceTickets = 50         // constant
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to our conferece %v booking app\n", conferenceName)
	fmt.Printf("Total No of Tickets: %v\nAvailable No of Tickets: %v\n", conferenceTickets, remainingTickets)
}
