package main

import (
	"fmt"
	"booking-app/helper"
	"time"
	"sync"
)

// package level variables cant be declared using :=
var conferenceName string = "Go Conference 2023"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // Array of 50 strings

// in case of running the program only once and not in a for loop, we can use waitgroups to wait for the goroutines to finish
var wg = sync.WaitGroup{}

type UserData struct {
	firstname string
	lastname string
	email string
	numberOfTickets uint
}

func main() {
	// %T is used to print the type of the variable
	// fmt.Printf("Conference name is %T, conference Tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstname, lastname, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookings = bookTicket(firstname, lastname, email, userTickets)

			wg.Add(1)
			go sendTicket(userTickets, firstname, lastname, email)
			
			firstNames := getFirstNames()
			fmt.Println("Bookings so far: ", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if(noTicketsRemaining) {
				fmt.Println("Sorry, all tickets are sold out!")
				break
			}
		}	else {
				if !isValidName {
					fmt.Println("Please enter a valid name")
				} 
				if !isValidEmail {
					fmt.Println("Please enter a valid email address")
				} 
				if !isValidTicketNumber {
					fmt.Printf("Oops! We have only %v tickets in lot!\n", remainingTickets)
					break
				} 
			}
	}
	wg.Wait()
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
				firstNames = append(firstNames, booking.firstname)
		}
	return firstNames
}

func getUserInput () (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastname)
	
	fmt.Print("Please enter your email address: ")
	fmt.Scan(&email)
	
	fmt.Print("Enter how many tickets you want to buy: ")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func bookTicket (firstname string, lastname string, email string, userTickets uint) ([]UserData){
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	var userData = UserData{
		firstname: firstname,
		lastname: lastname,
		email: email,
		numberOfTickets: userTickets,
	}

	// userData["firstname"] = firstname
	// userData["lastname"] = lastname
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	
	bookings = append(bookings, userData)
	fmt.Println("Bookings so far: ", bookings)

	fmt.Printf("Thank you %v %v, you have booked %v tickets for %v. You will receive a confirmation on the email address %v\n", firstname, lastname, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)

	return bookings
}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v", userTickets, firstname, lastname)
	fmt.Println("#################")
	fmt.Println("Sending", ticket, "to email address\n", email)
	fmt.Println("#################")
	wg.Done()
}
