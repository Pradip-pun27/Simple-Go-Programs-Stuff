package main

import (
	"fmt"
	"strings"
	"time"
)

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("___________________________________________")
	fmt.Printf("Sending Ticket:\n %v \nto email Address %v\n", Ticket, email)
	fmt.Println("___________________________________________")
	// wg.Done()
}

func GreetUser(ConfrncName string, Conf_Tkts int, remain_tkt uint) {
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Printf("Wlcm to %v booking application\n", ConfrncName)
	fmt.Printf("We've total of %v Tickets and %v are still a/v.\n", Conf_Tkts, remain_tkt)
	fmt.Println("Get yr Tickets here to attend")
	fmt.Println("-----------------------------------------------------------------------------------------")
}

func getBookingName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateInput(firstName string, lastName string, email string, userTickets uint, remain_tkt uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTktNumber := userTickets > 0 && userTickets <= remain_tkt
	return isValidEmail, isValidName, isValidTktNumber

}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Gmail:")
	fmt.Scan(&email)

	fmt.Println("Enter the no. of Tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(remain_tkt uint, bookings []string, userTickets uint, ConfrncName string, firstName string, lastName string, email string) {
	remain_tkt -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("TYVM %v %v for booking %v Tickets. You'll receive a Confirmation email to at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets left for %v\n\n\n", remain_tkt, ConfrncName)
	firstNames := getBookingName(bookings)
	fmt.Printf("These all are our Costumers' name: %v\n\n\n", firstNames) // Call PrintBookingName function
}

// var wg = sync.WaitGroup{}

func main() {
	ConfrncName := "Go Seminar"
	const Conf_Tkts int = 50
	var remain_tkt uint = 50
	var bookings []string
	GreetUser(ConfrncName, Conf_Tkts, remain_tkt) // Call greet user Function

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTktNumber := validateInput(firstName, lastName, email, userTickets, remain_tkt)
		if isValidName && isValidEmail && isValidTktNumber {
			bookTicket(remain_tkt, bookings, userTickets, ConfrncName, firstName, lastName, email)
			// wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			if remain_tkt == 0 {
				//exit the Program
				fmt.Println("All Reserved Sorry")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Given Name is Too Short")
			}
			if !isValidEmail {
				fmt.Println("Email Address u given doen't contain @ symbol.")
			}
			if !isValidTktNumber {
				fmt.Println("Number of Tks u given is Invalid")
			}
		}
	}
	// wg.Wait()

}
