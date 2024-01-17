package handlers

import (
	"encoding/json"
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/config"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
	"joe-l-mathew/Tambula-Ticket-Generator/models"
	"net/http"
)

func GenerateTicket(w http.ResponseWriter, r *http.Request) {
	generatedTicket := functions.GenerateTicket()
	var tickets []models.Ticket

	for _, val := range generatedTicket {
		stringMatrix, err := functions.MatrixToString(val)
		if err != nil {
			// Handle the error, for example, log it
			fmt.Println("Error converting matrix to string:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tickets = append(tickets, models.Ticket{MatrixData: stringMatrix})
	}

	// Add tickets to the database
	game := models.Game{
		Tickets: tickets,
	}

	if err := config.DB.Save(&game).Error; err != nil {
		// Handle the error, for example, log it
		fmt.Println("Error saving game to the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"ticket": game,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle the error, for example, log it
		fmt.Println("Error encoding JSON response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
