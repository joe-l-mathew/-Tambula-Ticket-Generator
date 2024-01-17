package main

import (
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/config"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
)

func main() {
	config.ConnectDatabase()
	generatedTickets := functions.GenerateTicket()
	for _, v := range generatedTickets {
		for _, v2 := range v {
			fmt.Println(v2)
		}
		fmt.Println("*******************")
	}
}
