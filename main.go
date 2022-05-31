package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // := only for variables, syntax sugar
	const conferenceTickets = 50
	var remainingTickets uint = 50 // only positive numbers
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// You can use exclamation prefix to use the negative form of this variable
		// !isValidCity is equal to city != "Singapore" && "London"
		// isValidCity := city == "Singapore" || "London"

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			// use _ as a blank identifier to not cause errors of not use variable
			// in this for loop, _ is the index of the list

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name are too short!")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && uint(userTickets) <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask for params
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName) // refer a pointer to store value with &

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)
	var firstNames []string = getFirstNames(bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}
