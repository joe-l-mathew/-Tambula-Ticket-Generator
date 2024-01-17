package functions

import (
	"sort"
)

func SortResult(initialSlice [][][]int) {
	for i := 0; i < 9; i++ {
		for cardIndex, individualCard := range initialSlice {
			var nonZeroCount int = 0
			var num1 int
			var num2 int
			var num3 int
			var num1Index int
			var num2Index int
			var num3Index int
			for rowIndex, rowElement := range individualCard {
				if rowElement[i] != 0 {
					nonZeroCount++
					if nonZeroCount == 1 {
						num1Index = rowIndex
						num1 = rowElement[i]
					}
					if nonZeroCount == 2 {
						num2Index = rowIndex
						num2 = rowElement[i]
					}
					if nonZeroCount == 3 {
						num3Index = rowIndex
						num3 = rowElement[i]
					}
				}
			}
			if nonZeroCount == 2 && num1 > num2 {
				initialSlice[cardIndex][num1Index][i] = num2
				initialSlice[cardIndex][num2Index][i] = num1
			} else if nonZeroCount == 3 {
				toSort := []int{num1, num2, num3}
				sort.Ints(toSort)
				initialSlice[cardIndex][num1Index][i] = toSort[0]
				initialSlice[cardIndex][num2Index][i] = toSort[1]
				initialSlice[cardIndex][num3Index][i] = toSort[2]

			}
		}

	}
}
