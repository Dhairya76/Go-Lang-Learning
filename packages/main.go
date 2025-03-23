package main

import (
	"fmt"
	"strings"
	"GOLANG/packages/utils"
)

// * ----------------------------------- GO CONFERENCE BOOKING SYSTEM -----------------------------------
// * This program allows users to book tickets for a Go Conference
// TODO: Spilt the code in different files just like packages

// * Global Variables
var conference_name string = "Go Conference" // ? Name of the conference
var bookings []string                        // ? Slice to store booking details

func main() {
	// * Infinite loop to keep the booking system running
	for {
		// * Display welcome message
		greet_user()

		// * Collect user details
		var first_name, last_name, email, user_tickets = get_user_input()

		// * Validate user input
		var is_valid_name, is_valid_email, is_valid_ticket_count bool = utils.Validate_user_input(first_name, last_name, email, user_tickets)

		if is_valid_name && is_valid_email && is_valid_ticket_count {
			// * Process ticket booking
			bookings = book_tickets(user_tickets, first_name, last_name, email)

			// * Display the first names of attendees
			print_first_names()
			var first_names_fetched = get_first_names()
			fmt.Printf("List of first names of guests from return function: %v\n", first_names_fetched)

		} else {
			// ! Display appropriate error messages based on invalid inputs
			if !is_valid_name {
				fmt.Println("❌ You have entered an invalid name. Please use at least 2 characters for both first and last name.")
			}
			if !is_valid_email {
				fmt.Println("❌ You have entered an invalid email. Please include '@' in your email address.")
			}
			if !is_valid_ticket_count {
				fmt.Println("❌ The number of tickets entered is invalid. Please enter a valid number.")
			}

			// * Using multiple `if` statements instead of `else if` to check all conditions
		}

		// * Check if tickets are sold out
		// * We can also import variables as well
		if utils.Remaining_tickets == 0 {
			fmt.Println("🚀 Our conference is SOLD OUT! No more tickets available.")
			break
		}
	}
}

// * Function: greet_user
// ? Displays a welcome message to the user
func greet_user() {
	fmt.Println("👋 Welcome to the Go Conference Booking System!")
}

// * Function: print_first_names
// ? Extracts and prints the first names of all booked attendees
func print_first_names() {
	name := []string{} // Slice to store extracted first names

	for _, booking := range bookings {
		// * Splitting the full name into parts (first name and last name)
		var names = strings.Fields(booking)
		name = append(name, names[0]) // * Adding only the first name to the slice
	}

	// * Printing the extracted first names
	fmt.Printf("📝 List of first names of guests: %v\n", name)
}

// * Function: get_first_names
// ? Extracts and returns the first names of all booked attendees
func get_first_names() []string {
	name := []string{} // Slice to store extracted first names

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		name = append(name, names[0]) // * Storing only the first name
	}
	return name
}

// * Function: get_user_input
// ? Collects user input for name, email, and ticket count
func get_user_input() (string, string, string, uint) {
	var first_name, last_name, email string
	var user_tickets uint

	fmt.Println("✏️ Enter your first name:")
	fmt.Scan(&first_name)

	fmt.Println("✏️ Enter your last name:")
	fmt.Scan(&last_name)

	fmt.Println("📧 Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("🎟️ How many tickets do you wish to buy?")
	fmt.Scan(&user_tickets)

	return first_name, last_name, email, user_tickets
}

// * Function: book_tickets
// ? Books tickets for the user and updates the system
// @param user_tickets - Number of tickets the user wants to buy
// @param first_name - User's first name
// @param last_name - User's last name
// @param email - User's email address
// @return bookings - Updated slice of all bookings
func book_tickets(user_tickets uint, first_name string, last_name string, email string) []string {
	// * Deduct booked tickets from remaining tickets
	utils.Remaining_tickets -= user_tickets

	// * Store the user's full name in the bookings slice
	bookings = append(bookings, first_name+" "+last_name)

	// * Confirmation message
	fmt.Printf("✅ Hey %v %v, your booking of %v tickets for %v was successful! 🎉\n", first_name, last_name, user_tickets, conference_name)
	fmt.Printf("📩 Tickets have been emailed to %v.\n", email)
	fmt.Printf("🎟️ Remaining tickets available: %v\n", utils.Remaining_tickets)

	return bookings
}
