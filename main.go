package main

import (
	"GOLANG/first"
	"fmt"
)

func main() {
	var conference_name string = "Go Conference"
	const conference_tickets int = 50
	var remaining_tickets uint = 50
	// although if the value is passed at the time of declaration
	// Go complier will understand the type but it is good practice to declare the data type

	// next_conference_name := "Docker Conference"
	// another way to declare the variables in GO
	// But only for var not for const and you cannot declare the type

	fmt.Println("Welcome to our conference booking application")
	fmt.Println("Next conference scheduled is:", conference_name)
	fmt.Println("Get your conference tickets asap, total tickets are:", conference_tickets)
	fmt.Printf("Get your conference tickets asap, %v tickets are available \n", remaining_tickets)
	first.Hello()
	// fmt.Println("Next conference scheduled is:", next_conference_name)

	var first_name string
	var last_name string
	var email string
	var user_tickets uint
	// ask the user to enter name and number of tickets

	fmt.Println("Enter your first name:")
	fmt.Scan(&first_name) // Name entered was Dhairya
	// by just fmt.Scan(first_name) it didn't paused or allowed us to write
	// it just executed the program with first_name as null
	// to avoid this we need to add a pointer fmt.Scan(&first_name) something like this

	// The pointer basically knnows the memory address of the variable
	// fmt.Println(first_name)  // Name printed Dhairya
	// fmt.Println(&first_name) // Memory address 0xc0000240d0

	fmt.Println("Enter your last name:")
	fmt.Scan(&last_name)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets you wish to buy:")
	fmt.Scan(&user_tickets)

	// Now we will add the booking ticket logic

	// Here the error will be type mismatch as one is int and another one is uint
	// so that needs to be handled

	remaining_tickets = remaining_tickets - user_tickets
	fmt.Printf("Hey, %v %v your booking of  %v tickets was successful and tickets are mailed to %v \n", first_name, last_name, user_tickets, email)
	// It prints Hey, Dhairya Thakker your booking of  3 tickets was successful and tickets are mailed to dt@gmail.com
	fmt.Printf("This many tickets are now availabe for the conference %v \n", remaining_tickets)

	//---------------------------------------------------- Now lets move to ARRAYS and SLICES in GO --------------------------------------------------------------
	// In GO we need to mention the length of the array at the time of declaration followed by data type

	// var bookings = [50]string{"Thor", "Tony", "Steve"}

	// Alternate way to declare Array which we are gonna use

	var bookings [50]string

	// Adding new values
	bookings[0] = first_name + " " + last_name
	// bookings[1] = first_name + " " + last_name

	fmt.Printf("The whole array is: %v\n", bookings)
	fmt.Printf("The first element in array is: %v\n", bookings[0])
	fmt.Printf("The type of bookings is %T\n", bookings)
	fmt.Printf("The length of array is %v\n", len(bookings))

	// Slices helps us to define the a list that is more dynamic in size
	// Slice is an abstraction of array
	// Slices are more flexible and powerful

	// Declare a slice

	var first_slice []string

	first_slice = append(first_slice, first_name+" "+last_name)

	fmt.Printf("The whole slice is: %v\n", first_slice)
	fmt.Printf("The first element in slice is: %v\n", first_slice[0])
	fmt.Printf("The type of bookings is %T\n", first_slice)
	fmt.Printf("The length of slice is %v\n", len(first_slice))

}
