package functions

func SortResult(initialSlice [][][]int) {
	for i := 0; i < 9; i++ {
		for cardIndex, individualCard := range initialSlice {
			var nonZeroCount int = 0
			var num1 int
			var num2 int
			var num1Index int
			var num2Index int
			for rowIndex, rowElement := range individualCard {
				if rowElement[i] != 0 {
					nonZeroCount++
					if nonZeroCount == 1 {
						num1Index = rowIndex
						num1 = rowElement[i]
					} else {
						num2Index = rowIndex
						num2 = rowElement[i]
					}
				}
			}
			if nonZeroCount == 2 && num1 > num2 {
				initialSlice[cardIndex][num1Index][i] = num2
				initialSlice[cardIndex][num2Index][i] = num1
			}
		}

	}
}
