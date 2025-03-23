package utils

import "strings"

// * Function: validate_user_input
// ? Validates user input (name, email, and ticket count)
// @return (bool, bool, bool) - Returns whether each input is valid
var Remaining_tickets uint = 50              // ? Total available tickets
func Validate_user_input(first_name string, last_name string, email string, user_tickets uint) (bool, bool, bool) {
	var is_valid_name bool = len(first_name) >= 2 && len(last_name) >= 2
	var is_valid_email bool = strings.Contains(email, "@")
	var is_valid_ticket_count bool = user_tickets > 0 && user_tickets <= Remaining_tickets

	return is_valid_name, is_valid_email, is_valid_ticket_count
}

// * Now since we have changed the package we need to export this function
// * and in go it is pretty simple just make the first letter of the function capitalize
