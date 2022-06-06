package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// You can use exclamation prefix to use the negative form of this variable
	// !isValidCity is equal to city != "Singapore" && "London"
	// isValidCity := city == "Singapore" || "London"
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && uint(userTickets) <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
