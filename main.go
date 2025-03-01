package main

import (
	"GOLANG/first"
	"fmt"
)

func main() {
	var conference_name = "Go Conference"
	const conference_tickets = 50
	var remaining_tickets = 50
	fmt.Println("Welcome to our conference booking application")
	fmt.Println("Next conference scheduled is:", conference_name)
	fmt.Println("Get your conference tickets asap, total tickets are:", conference_tickets)
	fmt.Printf("Get your conference tickets asap, %v tickets are available \n", remaining_tickets)
	first.Hello()
}
