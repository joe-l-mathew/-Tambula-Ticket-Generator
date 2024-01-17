package functions

import (
	"fmt"
	"time"
)

func GenerateTicket() [][][]int {
	startTime := time.Now()
	cardCount, rows, columns, elementCount := 6, 3, 9, 90

	initialSlice := make([][][]int, cardCount)
	for i := range initialSlice {
		initialSlice[i] = make([][]int, rows)
		for j := range initialSlice[i] {
			initialSlice[i][j] = make([]int, columns)
		}
	}
	isPlacedSuccesfully := false

	for !isPlacedSuccesfully {
		isPlacedSuccesfully = PlaceElement(elementCount, cardCount, initialSlice)
		if !isPlacedSuccesfully {
			initialSlice = make([][][]int, cardCount)
			for i := range initialSlice {
				initialSlice[i] = make([][]int, rows)
				for j := range initialSlice[i] {
					initialSlice[i][j] = make([]int, columns)
				}
			}
		}
	}
	SortResult(initialSlice)
	fmt.Println(time.Since(startTime))
	return initialSlice
}
