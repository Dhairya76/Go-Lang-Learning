package main

import (
	"fmt"
	"strings"
)

//---------------------------------------------------- LOOPS IN GO --------------------------------------------------------------
//  * Note - In Go, loops are simplified since we only have the `for` loop.
//  * There are no `while`, `do-while`, or `for-each` loops like in other languages.

func main() {
	fmt.Println("We are learning loops in GO")

	// Conference details
	var conference_name string = "Go Conference"
	var remaining_tickets uint = 50 // Total available tickets
	var bookings []string           // Slice to store booking details

	// ----------------------------------- INFINITE LOOP -----------------------------------
	//  ! Warning: This loop will keep running until manually stopped (Ctrl + C) or program exit conditions are added.

	// TODO : Need to add conditional statements to end the loop
	for {
		// Variables to store user input
		var first_name string
		var last_name string
		var email string
		var user_tickets uint

		// Taking user input
		fmt.Println("Enter your first name:")
		fmt.Scan(&first_name)

		fmt.Println("Enter your last name:")
		fmt.Scan(&last_name)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("How many tickets you wish to buy:")
		fmt.Scan(&user_tickets)

		// Updating the number of remaining tickets
		remaining_tickets -= user_tickets

		// Confirming booking to the user
		fmt.Printf("Hey %v %v, your booking of %v tickets for %v was successful. Tickets have been mailed to %v.\n",
			first_name, last_name, user_tickets, conference_name, email)

		fmt.Printf("Remaining tickets available for the conference: %v\n", remaining_tickets)

		// Storing the user's full name in the bookings slice
		bookings = append(bookings, first_name+" "+last_name)

		// Displaying the list of guests
		fmt.Printf("Current list of guests: %v\n", bookings)

		// ----------------------------------- ITERATING OVER A LIST -----------------------------------
		// We will iterate over the `bookings` slice and extract only the first names of the users.

		name := []string{} // Slice to store extracted first names

		for _, booking := range bookings {
			// Splitting the full name into parts (first name and last name)
			var names = strings.Fields(booking)

			// Appending only the first name to the `name` slice
			name = append(name, names[0])

			// Printing the extracted first names
			fmt.Printf("List of first names of guests: %v\n", name)
		}
	}
}
