package main

import (
	"fmt"
	"joe-l-mathew/Tambula-Ticket-Generator/functions"
)

func main() {
	cardCount, rows, columns, elementCount := 6, 3, 9, 90
	// Create a 3D slice with dimensions cardCount x rows x columns
	initialSlice := make([][][]int, cardCount)
	for i := range initialSlice {
		initialSlice[i] = make([][]int, rows)
		for j := range initialSlice[i] {
			initialSlice[i][j] = make([]int, columns)
		}
	}
	//looping from number 1 to 90
	for i := 1; i < elementCount; i++ {
		valuePlaced := false
		for !valuePlaced {
			cardIndex := functions.GenerateRandomNumber(0, cardCount)
			columnIndex := functions.GetColumnIndex(i)
			rowIndex := functions.GenerateRandomNumber(0, 3)
			//check for zero
			if initialSlice[cardIndex][rowIndex][columnIndex] == 0 {
				//check for max 2 in column
				numberOfColumsFilled := 0
				numberOfRowElementsFilled := 0
				for _, rowSlice := range initialSlice[cardIndex] {
					if rowSlice[columnIndex] != 0 {
						numberOfColumsFilled++
					}
				}

				if numberOfColumsFilled < 2 {
					fmt.Println("true", i)
					for _, val := range initialSlice[cardIndex][rowIndex] {
						if val != 0 {
							numberOfRowElementsFilled++
						}

					}
					if numberOfRowElementsFilled < 5 {
						initialSlice[cardIndex][rowIndex][columnIndex] = i
						valuePlaced = true
					}

					//check weather the row has 5 elements filled

				}

			}

		}
		//1 will be recieved initially

	}

	for _, v := range initialSlice {
		for _, v2 := range v {
			fmt.Println(v2)
		}
		fmt.Println("*******************")
	}
}
