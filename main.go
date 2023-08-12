package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference 2023"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// %T is used to print the type of the variable
	// fmt.Printf("Conference name is %T, conference Tickets is %T, remaining tickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to", conferenceName, "booking application")
	// Printf is used to format the string
	fmt.Printf("There are total %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets now!")

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstname)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastname)

	fmt.Println("Enter how many tickets you want to buy")
	fmt.Scan(&userTickets)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)

	remainingTickets = conferenceTickets - userTickets

	fmt.Printf("Thank you %v %v, you have booked %v tickets for %v. You will receive a confirmation on the email address %v\n", firstname, lastname, userTickets, conferenceName, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)

}