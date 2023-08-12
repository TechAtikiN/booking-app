package main

import (
	"fmt"
	"strings"
)

// package level variables cant be declared using :=
	var conferenceName string = "Go Conference 2023"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string // Array of 50 strings

func main() {
	// %T is used to print the type of the variable
	// fmt.Printf("Conference name is %T, conference Tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstname, lastname, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := validateUserInput(firstname, lastname, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookings = bookTicket(firstname, lastname, email, userTickets)
			
			firstNames := getFirstNames()
			fmt.Println("Bookings so far: ", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if(noTicketsRemaining) {
				fmt.Println("Sorry, all tickets are sold out!")
				break
			}	
		}	else{
				if !isValidName {
					fmt.Println("Please enter a valid name")
				} 
				if !isValidEmail {
					fmt.Println("Please enter a valid email address")
				} 
				if !isValidTicketNumber {
					fmt.Println("Please enter a valid number of tickets")
				} 
			}
	}
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	// Printf is used to format the string
	fmt.Printf("There are total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
		}
	return firstNames
}

func validateUserInput(firstname string,
	lastname string,
	email string,
	userTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >=2 && len(lastname) >=2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput () (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstname)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastname)
	
	fmt.Println("Please enter your email address")
	fmt.Scan(&email)
	
	fmt.Println("Enter how many tickets you want to buy")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTicket (firstname string, lastname string, email string, userTickets uint) ([]string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstname + " " + lastname)

	fmt.Printf("Thank you %v %v, you have booked %v tickets for %v. You will receive a confirmation on the email address %v\n", firstname, lastname, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)
	return bookings
}