package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber

}
