package main

import (
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
	"time"
)

func main() {
	startTime := time.Now()
	cardCount, rows, columns, elementCount := 6, 3, 9, 90
	// Create a 3D slice with dimensions cardCount x rows x columns
	initialSlice := make([][][]int, cardCount)
	for i := range initialSlice {
		initialSlice[i] = make([][]int, rows)
		for j := range initialSlice[i] {
			initialSlice[i][j] = make([]int, columns)
		}
	}
	isPlacedSuccesfully := false
	//looping from number 1 to 90
	//check for zero
	//check for max 2 in column
	//check weather the row has 5 elements filled
	//1 will be recieved initially
	for !isPlacedSuccesfully {
		isPlacedSuccesfully = functions.PlaceElement(elementCount, cardCount, initialSlice)
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

	for _, v := range initialSlice {
		for _, v2 := range v {
			fmt.Println(v2)
		}
		fmt.Println("*******************")
	}

	fmt.Println(time.Since(startTime))
}
