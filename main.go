package main

import (
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
)

func main() {
	//##STEPS TO FOLLOW#//
	// Create a 3D slice with dimensions cardCount x rows x columns
	//looping from number 1 to 90
	//check for zero
	//check for max 2 in column
	//check weather the row has 5 elements filled
	//1 will be recieved initially
	generatedTickets := functions.GenerateTicket()
	for _, v := range generatedTickets {
		for _, v2 := range v {
			fmt.Println(v2)
		}
		fmt.Println("*******************")
	}
}
