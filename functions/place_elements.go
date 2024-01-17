package functions

func PlaceElement(elementCount int, cardCount int, initialSlice [][][]int) bool {

	for i := 1; i <= elementCount; i++ {
		iterationCount := 0
		valuePlaced := false
		for !valuePlaced {
			cardIndex := GenerateRandomNumber(0, cardCount)
			columnIndex := GetColumnIndex(i)
			rowIndex := GenerateRandomNumber(0, 3)
			if iterationCount == 1000 {
				return false
			}

			if initialSlice[cardIndex][rowIndex][columnIndex] == 0 {

				numberOfColumsFilled := 0
				numberOfRowElementsFilled := 0
				for _, rowSlice := range initialSlice[cardIndex] {
					if rowSlice[columnIndex] != 0 {
						numberOfColumsFilled++
					}
				}

				for _, val := range initialSlice[cardIndex][rowIndex] {
					if val != 0 {
						numberOfRowElementsFilled++
					}

				}
				if numberOfRowElementsFilled < 5 && (numberOfRowElementsFilled <= (columnIndex-1) || columnIndex < 2) {

					initialSlice[cardIndex][rowIndex][columnIndex] = i
					valuePlaced = true

				} else {
					iterationCount++
				}

			}

		}

	}
	for count := 0; count < 9; count++ {
		for _, individualCard := range initialSlice {
			var zeroCount int = 0
			for _, rowElement := range individualCard {
				if rowElement[count] == 0 {
					zeroCount++
				}
			}
			if zeroCount == 3 {
				return false
			}

		}
	}

	return true
}
