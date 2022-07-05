package main

import (
	"fmt"
	"strings"
	"time"
)

const totalTickets uint = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = totalTickets
	bookings              = make([]UserData, 0) // empty slice of struct
)

type UserData struct {
	firstName string
	lastName  string
	email     string
	ticketNo  uint
}

func main() {

	fmt.Println("***********************************************")

	greetUser()

	var flag bool = true
	for flag && remainingTickets > 0 {

		firstName, lastName, email, bookedTickets := GetuserDetails()
		//userInput validation
		isCorrectName, isCorrectEmail, isCorrectTickets := validateUser(firstName, lastName, email, bookedTickets)

		if isCorrectName && isCorrectEmail && isCorrectTickets {

			bookTicket(firstName, lastName, email, bookedTickets)
			fmt.Println("Remaining tickets are : ", remainingTickets)

			go sendTickets(firstName, bookedTickets, email, conferenceName)

			BookedFirstName()

			fmt.Println("WANT TO BOOK For Another User??  [Y/N]")
			var again string
			fmt.Scan(&again)
			if again == "Y" || again == "y" {
				flag = true
			} else {
				flag = false
			}

		} else {
			fmt.Printf("Hey! Data You entered in not Correct\n")
			if !isCorrectName {
				fmt.Println("Name is too short")
			}
			if !isCorrectEmail {
				fmt.Println("Email must have '@' in it")
			}
			if !isCorrectTickets {
				fmt.Printf("we have only %v tickets left, please enter valid input/n", remainingTickets)
			}
			continue
		}
		fmt.Println("***********************************************")
	}
}

//-********************************************************************************************

func greetUser() {
	fmt.Println()
	fmt.Printf("Welcome to the %v system", conferenceName)
	fmt.Println("Get your tickets booking here")
}

func GetuserDetails() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var bookedTickets uint

	fmt.Println("Enter First Name ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last Name ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email ID ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of tickets needed ")
	fmt.Scan(&bookedTickets)

	return firstName, lastName, email, bookedTickets
}

func validateUser(firstName string, lastName string, email string, bookedTickets uint) (bool, bool, bool) {
	isCorrectName := len(firstName) >= 2 && len(lastName) >= 2
	isCorrectEmail := strings.Contains(email, "@")
	isCorrectTickets := bookedTickets <= remainingTickets

	return isCorrectName, isCorrectEmail, isCorrectTickets
}

func bookTicket(firstName string, lastName string, email string, bookedTickets uint) /*uint*/ {

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		ticketNo:  bookedTickets,
	}

	bookings = append(bookings, userData) //adding new boooking in slice of struct

	fmt.Printf("user info is %v\n", userData)
	fmt.Printf("Hey! %v %v , Thanks for booking %v Tickets, \nYou'll get confirmation mail on %v \n", firstName, lastName, bookedTickets, email)

	remainingTickets = remainingTickets - bookedTickets
	//return remainingTickets

}

func BookedFirstName() {
	person := []string{}
	for _, value := range bookings {
		person = append(person, value.firstName)

	}
	fmt.Printf("\nTickets booked by : %v\n", person)
}

func sendTickets(firstName string, bookedTickets uint, email string, confName string) {
	time.Sleep(50 * time.Second)
	fmt.Println("*********************************")
	fmt.Printf("Hi, %v , thanks for booking %v tickets,\nTicket are being send to %v,\n See you at %v", firstName, bookedTickets, email, confName)
	fmt.Println("*********************************")

}
