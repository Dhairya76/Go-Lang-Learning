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
	next_conference_name  := "Docker Conference"
	// another way to declare the variables in GO
	// But only for var not for const and you cannot declare the type


	fmt.Println("Welcome to our conference booking application")
	fmt.Println("Next conference scheduled is:", conference_name)
	fmt.Println("Get your conference tickets asap, total tickets are:", conference_tickets)
	fmt.Printf("Get your conference tickets asap, %v tickets are available \n", remaining_tickets)
	first.Hello()
	fmt.Println("Next conference scheduled is:", next_conference_name)


	var user_name string
	var user_tickets int
	// ask the user to enter name and number of tickets

	user_name = "Steve Rogers"
	user_tickets = 3
	fmt.Println(user_name)
	fmt.Println(user_tickets)
	fmt.Printf("The User,%v booked %v tickets \n", user_name, user_tickets)
	fmt.Printf("The type of user_name value is %T and type of user_tickets is %T \n", user_name, user_tickets)

}
