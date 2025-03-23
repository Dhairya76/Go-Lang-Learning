package main

import (
	"GOLANG/goroutines-concurrency/utils"
	"fmt"
	"sync"
	"time"
)

// * ----------------------------------- GO CONFERENCE BOOKING SYSTEM -----------------------------------
// * This program allows users to book tickets for a Go Conference

// * Global Variables
var conference_name string = "Go Conference" // * Name of the conference
var bookings = make([]user_data, 1)          // * Slice to store booking details

// * Struct to hold user booking details
type user_data struct {
	first_name     string
	last_name      string
	email          string
	tickets_booked uint
}

var wg = sync.WaitGroup{} // * WaitGroup to manage concurrency

func main() {
	// * Infinite loop to keep the booking system running

	// ! If we remove this loop and allow just one ticket to be booked, 
	// ! the main thread will end immediately without waiting for send_tickets to complete.
	// ? To prevent this, we use WaitGroup to make the main thread wait.

	// * Display welcome message
	greet_user()

	// * Collect user details
	var first_name, last_name, email, user_tickets = get_user_input()

	// * Validate user input
	var is_valid_name, is_valid_email, is_valid_ticket_count bool = utils.Validate_user_input(first_name, last_name, email, user_tickets)

	if is_valid_name && is_valid_email && is_valid_ticket_count {
		// * Process ticket booking
		bookings = book_tickets(user_tickets, first_name, last_name, email)

		// ! send_tickets has a 10-second delay, which blocks execution.
		// ! This means other users cannot book tickets until the current one completes.
		// ? To fix this, we run send_tickets as a goroutine.

		wg.Add(1)                         // * Adding a goroutine to WaitGroup
		go send_tickets(user_tickets, first_name, last_name, email)

		// * Display the first names of attendees
		print_first_names()
		var first_names_fetched = get_first_names()
		fmt.Printf("List of first names of guests from return function: %v\n", first_names_fetched)

	} else {
		// ! Display appropriate error messages based on invalid inputs
		if !is_valid_name {
			fmt.Println("❌ Invalid name! Please use at least 2 characters for both first and last name.")
		}
		if !is_valid_email {
			fmt.Println("❌ Invalid email! Please include '@' in your email address.")
		}
		if !is_valid_ticket_count {
			fmt.Println("❌ Invalid ticket count! Please enter a valid number.")
		}
	}

	wg.Wait() // * Wait for all goroutines to finish

	// * Check if tickets are sold out
	if utils.Remaining_tickets == 0 {
		fmt.Println("🚀 Our conference is SOLD OUT! No more tickets available.")
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
	name := []string{} // * Slice to store extracted first names

	for _, booking := range bookings {
		name = append(name, booking.first_name) // * Adding only the first name to the slice
	}

	// * Printing the extracted first names
	fmt.Printf("📝 List of first names of guests: %v\n", name)
}

// * Function: get_first_names
// ? Extracts and returns the first names of all booked attendees
func get_first_names() []string {
	name := []string{} // * Slice to store extracted first names

	for _, booking := range bookings {
		name = append(name, booking.first_name) // * Storing only the first name
	}
	return name
}

// * Function: get_user_input
// ? Collects user input for name, email, and ticket count
// @return first_name, last_name, email, user_tickets
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
// @param user_tickets (uint) - Number of tickets to book
// @param first_name (string) - First name of the user
// @param last_name (string) - Last name of the user
// @param email (string) - Email address of the user
// @return updated bookings slice
func book_tickets(user_tickets uint, first_name string, last_name string, email string) []user_data {
	// * Deduct booked tickets from remaining tickets
	utils.Remaining_tickets -= user_tickets

	// * Creating a struct for user data
	var user_data = user_data{
		first_name:     first_name,
		last_name:      last_name,
		email:          email,
		tickets_booked: user_tickets,
	}

	// * Store the user's full name in the bookings slice
	bookings = append(bookings, user_data)
	fmt.Println("Bookings:", bookings)

	// * Confirmation message
	fmt.Printf("✅ Hey %v %v, your booking of %v tickets for %v was successful! 🎉\n", first_name, last_name, user_tickets, conference_name)
	fmt.Printf("📩 Tickets have been emailed to %v.\n", email)
	fmt.Printf("🎟️ Remaining tickets available: %v\n", utils.Remaining_tickets)

	return bookings
}

// * Function: send_tickets
// ? Simulates sending tickets via email (includes a delay for realism)
// @param user_tickets (uint) - Number of tickets booked
// @param first_name (string) - First name of the user
// @param last_name (string) - Last name of the user
// @param email (string) - Email address of the user
func send_tickets(user_tickets uint, first_name string, last_name string, email string) {
	// ? In real-world applications, generating tickets takes time.
	// ? To simulate this, we add a delay.

	time.Sleep(10 * time.Second) // * Simulated ticket processing delay

	// * Formatting ticket details
	var tickets = fmt.Sprintf("%v tickets for %v %v", user_tickets, first_name, last_name)

	// * Sending ticket confirmation
	fmt.Println("------------------------")
	fmt.Printf("📩 Sending ticket:\n %v to email address %v\n", tickets, email)
	fmt.Println("------------------------")

	wg.Done() // * Mark goroutine as done
}
