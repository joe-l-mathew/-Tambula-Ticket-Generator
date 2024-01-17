package handlers

import (
	"encoding/json"
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/config"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
	"joe-l-mathew/Tambula-Ticket-Generator/models"
	"net/http"
	"sort"
	"strconv"
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

	// Sort tickets by TicketID
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i].ID < tickets[j].ID
	})

	// Add tickets to the database
	game := models.Game{
		Tickets: tickets,
	}

	if err := config.DB.Save(&game).Error; err != nil {
		fmt.Println("Error saving game to the database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"tickets": make(map[string][][]int),
	}

	for _, ticket := range tickets {
		matrix, _ := functions.StringToMatrix(ticket.MatrixData)
		response["tickets"].(map[string][][]int)[fmt.Sprintf("%d", ticket.ID)] = matrix
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("Error encoding json response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	itemsPerPageStr := r.URL.Query().Get("itemsPerPage")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	itemsPerPage, err := strconv.Atoi(itemsPerPageStr)
	if err != nil || itemsPerPage < 1 {
		itemsPerPage = 10
	}

	offset := (page - 1) * itemsPerPage
	var games []models.Game
	result := config.DB.Preload("Tickets").Limit(itemsPerPage).Offset(offset).Find(&games)
	if result.Error != nil {
		fmt.Println("Error retrieving games with tickets:", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"games": games,
		"page":  page,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("Error encoding JSON response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
