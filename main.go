package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceName, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user their name
	// fmt.Scan(&userName)

	// variable goes to memory address
	// fmt.Println(remainingTickets)
	// Example using pointer :
	// fmt.Println(&remainingTickets)
	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
