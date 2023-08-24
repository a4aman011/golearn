package main

import (
	"fmt"
	"strings"
)

func main() {
	cfrcName := "Go Conference"
	cfrcTicket := 50
	var remTicket uint = 50
	bookings := []string{}

	greetUsers(cfrcName, cfrcTicket, remTicket)

	for {
		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remTicket)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remTicket, userTicket, bookings, firstName, lastName, email, cfrcName)

			//fmt.Printf("The whole slice: %v\n", bookings)
			//fmt.Printf("The first value: %v\n", bookings[0])
			//fmt.Printf("slice type: %T\n", bookings)
			//fmt.Printf("slice length: %v\n", len(bookings))

			firstNames := getFirstNames(bookings)
			fmt.Printf("These are all our bookings %v\n", firstNames)

			if remTicket == 0 {
				//end program
				fmt.Println("We are SOLD OUT!!. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
		}

	}
}

func greetUsers(cfrcName string, cfrcTicket int, remTicket uint) {
	fmt.Printf("Welcome to %v booking application\n", cfrcName)
	fmt.Printf("We have %v tickets in total and %v tickets still available\n", cfrcTicket, remTicket)
	fmt.Println("Get your tickets here")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTicket uint, remTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remTicket
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func bookTicket(remTicket uint, userTicket uint, bookings []string, firstName string, lastName string, email string, cfrcName string) {
	remTicket = remTicket - userTicket
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v ticket, you will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remTicket, cfrcName)

}
