package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "go conference" // variable name must be used
	const conferenceTickets uint = 50           // constant
	var remainingTickets int = 50
	// var bookings [50]string // defining an array var
	var bookings []string // defining an slice var

	fmt.Printf("Welcome to our conferece %v booking app\n", conferenceName)
	fmt.Printf("Total No of Tickets: %v\n", conferenceTickets)

	for {
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

		// if numTcks > remainingTickets {
		// 	fmt.Printf("We have only %v tickets remaining , you can't book %v tickets.\n", remainingTickets, numTcks)
		// 	continue
		// }

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && len(email) >= 3
		isValidTcksNum := numTcks <= remainingTickets && numTcks > 0

		if isValidName && isValidEmail && isValidTcksNum {

			fmt.Printf("congratulations , %s %s you have %v tickets.Please check %s for more detail\n", firstName, lastName, numTcks, email)
			remainingTickets = remainingTickets - numTcks

			// bookings[0] = firstName + " " + lastName //when array
			// bookings[4] = firstName + " " + lastName //when array

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Length booking array: %v\n", len(bookings))

			fmt.Printf("Remaining No of Tickets: %v\n", remainingTickets)

			//printing first name
			var firstNames []string

			for _, booking := range bookings {
				var name = strings.Fields(booking) // like python split on space
				firstNames = append(firstNames, name[0])
			}

			fmt.Printf("First Name of ticket bookers : %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("All Tickets sold out, come back next yeat")
				break
			}
		} else {

			if !isValidName {
				fmt.Printf("first name or last name is invalid entered . \n")
			}
			if !isValidEmail {
				fmt.Printf("Email address is not correct format\n")
			}
			if !isValidTcksNum {
				fmt.Printf("We have %v tickets remaining , you can't book %v tickets.\n", remainingTickets, numTcks)
			}
		}
	}
}
