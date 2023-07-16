// Go Organized into packages
// Packages are folders that contains .go files of yours

package main

import (
	"booking-app/shared"
	"fmt"
	"time"
)

// what we should have as value, and what we should have for these value to be stored? in variables
const eventTickets = 50

var eventName = "Go Event"
var remainingTickets = 50

// Who is attending the event
// Who booked the tickets?
// we can do that by creating array
// Also there is a "Slice" which is an abstraction of an Array to get array more dynamic flexible resized when needed
// but we user struct. Not map sorry map is boring
var bookings = make([]UserData, 0)

type UserData struct {
	userName    string
	email       string
	userTickets int
}

func main() {

	//calling greetuser function
	greetUsers()

	// In real life scenario this would be web app with a UI and a database connected to it
	//Also we want to create a constant loop
	// so App should not exit

	// IMPORTANT >>> IN goLang There is only one loop that is "FOR" we do not have (while)- (do-while) or (for-each)

	for remainingTickets > 0 && len(bookings) < 50 {

		userName, email, userTickets := getUserInput()

		//making sure client is not gonna buy more than 50 by if else block

		isValidName, isvalidEmail, isvalidTicketNumber := shared.ValidateUserInput(userName, email, userTickets, remainingTickets)

		if isValidName && isvalidEmail && isvalidTicketNumber {

			bookTicket(remainingTickets, userTickets, userName, email, eventName)
			go sendTicket(userTickets, userName, email)

			noTicketsRemaining := remainingTickets == 0

			//if-else check whether all ticket sold or not
			if noTicketsRemaining {
				fmt.Println("All tickets sold! Come back Next week")
				break
			}
		} else {
			fmt.Printf("Your input data is invalid. Try Again!")

		}

	}

	// In real life scenario this would be web app with a UI and a database connected to it
	//Also we want to create a constant loop
	// so App should not exit
	//Loop statemnt allows us to execute code multiple times, in a loop

}

func greetUsers() {

	fmt.Printf("\tWelcome To %v !!!\n We hope you enjoy your time here\n", eventName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", eventTickets, remainingTickets)
	fmt.Printf("Book a ticket right now in here!")

}

func getUserInput() (string, string, int) {

	//This is user enters:

	//explicitly defining type also works
	var userName string
	var email string
	var userTickets int

	// simple pointers used to get memory location of value with '&' mark
	fmt.Print("\nPlease enter your username: ") //ask for user's input using fmt
	fmt.Scan(&userName)
	fmt.Print("\nPlease enter your mail adress: ") //ask for user's email using fmt
	fmt.Scan(&email)
	fmt.Print("\nPlease enter number of tickets: ") //ask for user's input using fmt
	fmt.Scan(&userTickets)

	return userName, email, userTickets
}

func bookTicket(remainingTickets int, userTickets int, userName string, email string, eventName string) {

	var userData = UserData{
		userName:    userName,
		email:       email,
		userTickets: userTickets,
	}
	// simple logic to get remaining tickets
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, userData)

	fmt.Printf("Thank You %v for booking %v tickets from us! Your confirmation mail will be sent to your  %v adress\n", userName, userTickets, email)

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}

// Concurreny to make it multiple threads in program
func sendTicket(userTickets int, userName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Printf("##########")
	fmt.Printf("Sending ticket\n %v\n to email adress %v", ticket, email)

}
