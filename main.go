package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // := only for variables, syntax sugar
	const conferenceTickets = 50
	var remainingTickets uint = 50 // only positive numbers

	fmt.Printf("conferenceName is %T, remainingTickets is %T, conferenceTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	firstNames := []string{}
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		var bookings []string

		// ask for params
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName) // refer a pointer to store value with &

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// use _ as a blank identifier to not cause errors of not use variable
		// in this for loop, _ is the index of the list
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}
