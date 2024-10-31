package main

import (
	"fmt"
	"time"
	"sync"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookTicket(firstName, lastName, email, userTickets)

			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, email)

			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			if remainingTickets == 0{
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else{	
			if !isValidName{
				fmt.Println("First Name or Last Name needs to be at least 2 characters long")
			}
			if !isValidEmail{
				fmt.Println("Email is invalid")
			}
			if !isValidTicketNumber{
				fmt.Println("Invalid number of tickets. Please enter a number between 1 and", remainingTickets)
			}
		}

	}
	wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint){
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scanln(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scanln(&email)
	fmt.Println("Please enter the number of tickets you would like to book")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint){
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,	
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, userTickets uint, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("############################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############################################")
	wg.Done()
}