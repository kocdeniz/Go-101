// A function in shared, to be explicitly accessible for main function

// we simply uppercase the function name's first character
package shared

import "strings"

func ValidateUserInput(userName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {

	//user input validation plus
	var isValidName bool = len(userName) >= 2 && len(email) >= 2
	var isvalidEmail = strings.Contains(email, "@")
	var isvalidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isvalidEmail, isvalidTicketNumber

}
