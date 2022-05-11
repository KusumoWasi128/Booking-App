package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var wg sync.WaitGroup

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {

	greetUsers()

	var firstName string
	var lastName string
	var email string
	var userTickets int
	for {
		firstName, lastName, email, userTickets = getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames()
			fmt.Printf("the first names of bookings are: %v\n", firstNames)
			wg.Wait()
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	//	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avalaible\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//remainingTickets -= uint(userTickets)
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(12 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket %v to email address %v", ticket, email)
	fmt.Println("################")
	wg.Done()
}
