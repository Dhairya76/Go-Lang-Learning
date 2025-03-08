package main

import (
	"fmt"
	"strings"
)

// * ----------------------------------- IF LOOP -----------------------------------

func main() {
	fmt.Println("We are IF Else loops in GO")

	// Conference details
	var conference_name string = "Go Conference"
	var remaining_tickets uint = 50 // Total available tickets
	var bookings []string           // Slice to store booking details

	// for {
	// 	// Variables to store user input
	// 	var first_name string
	// 	var last_name string
	// 	var email string
	// 	var user_tickets uint

	// 	// Taking user input
	// 	fmt.Println("Enter your first name:")
	// 	fmt.Scan(&first_name)

	// 	fmt.Println("Enter your last name:")
	// 	fmt.Scan(&last_name)

	// 	fmt.Println("Enter your email address:")
	// 	fmt.Scan(&email)

	// 	fmt.Println("How many tickets you wish to buy:")
	// 	fmt.Scan(&user_tickets)

	// 	if user_tickets > remaining_tickets {
	// 		fmt.Println("Total tickets remaining are:", remaining_tickets, "So please book your tickets accordingly")
	// 		// break // * Here we exit the application if user exceeds the limit , but instead of break we can add continue to go to next iteration
	// 		continue
	// 	}

	// 	// Updating the number of remaining tickets
	// 	remaining_tickets -= user_tickets

	// 	// Confirming booking to the user
	// 	fmt.Printf("Hey %v %v, your booking of %v tickets for %v was successful. Tickets have been mailed to %v.\n",
	// 		first_name, last_name, user_tickets, conference_name, email)

	// 	fmt.Printf("Remaining tickets available for the conference: %v\n", remaining_tickets)

	// 	// Storing the user's full name in the bookings slice
	// 	bookings = append(bookings, first_name+" "+last_name)

	// 	// Displaying the list of guests
	// 	fmt.Printf("Current list of guests: %v\n", bookings)

	// 	// ----------------------------------- ITERATING OVER A LIST -----------------------------------
	// 	// We will iterate over the `bookings` slice and extract only the first names of the users.

	// 	name := []string{} // Slice to store extracted first names

	// 	for _, booking := range bookings {
	// 		// Splitting the full name into parts (first name and last name)
	// 		var names = strings.Fields(booking)

	// 		// Appending only the first name to the `name` slice
	// 		name = append(name, names[0])

	// 		// Printing the extracted first names
	// 		fmt.Printf("List of first names of guests: %v\n", name)
	// 	}

	// 	if remaining_tickets == 0 {
	// 		fmt.Println("Our conference is sold out")
	// 		break
	// 	}

	// 	// ? Nice our code ran well when user booked all the tickets
	// 	// ? but what will happen if the first user tries to book more than 50 tickets , let's see

	// 	// * It prints the following
	// 	// * Hey d t, your booking of 56 tickets for Go Conference was successful. Tickets have been mailed to d@gmail.com.
	// 	// * Remaining tickets available for the conference: 18446744073709551610
	// 	// * Current list of guests: [d t]
	// 	// * List of first names of guests: [d]
	// 	// * Enter your first name:
	// 	// * It didn't stop because the remaining tickets value is now some weird number

	// 	// ! So now we need to handle this

	// 	// We will add this check where we ask the user how many tickets they wanna book

	// 	// * After adding the check above it shows
	// 	// * Enter your first name:
	// 	// * d
	// 	// * Enter your last name:
	// 	// * t
	// 	// * Enter your email address:
	// 	// * d@gmail.com
	// 	// * How many tickets you wish to buy:
	// 	// * 56
	// 	// * We only have 50 tickets remaining
	// }

	// Let's Refactor the above code using IF ELSE

	// * ----------------------------------- IF ELSE LOOP -----------------------------------

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

		if user_tickets <= remaining_tickets {
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

		} else {
			fmt.Println("Total tickets remaining are:", remaining_tickets, "So please book your tickets accordingly")

		}

		if remaining_tickets == 0 {
			fmt.Println("Our conference is sold out")
			break
		}

		// ? Nice our code ran well when user booked all the tickets
		// ? but what will happen if the first user tries to book more than 50 tickets , let's see

		// * It prints the following
		// * Hey d t, your booking of 56 tickets for Go Conference was successful. Tickets have been mailed to d@gmail.com.
		// * Remaining tickets available for the conference: 18446744073709551610
		// * Current list of guests: [d t]
		// * List of first names of guests: [d]
		// * Enter your first name:
		// * It didn't stop because the remaining tickets value is now some weird number

		// ! So now we need to handle this

		// We will add this check where we ask the user how many tickets they wanna book

		// * After adding the check above it shows
		// * Enter your first name:
		// * d
		// * Enter your last name:
		// * t
		// * Enter your email address:
		// * d@gmail.com
		// * How many tickets you wish to buy:
		// * 56
		// * We only have 50 tickets remaining
	}

}
