package functions

func GetColumnIndex(value int) int {
	if value < 10 {
		return 0
	} else if value < 20 {
		return 1
	} else if value < 30 {
		return 2
	} else if value < 40 {
		return 3
	} else if value < 50 {
		return 4
	} else if value < 60 {
		return 5
	} else if value < 70 {
		return 6
	} else if value < 80 {
		return 7
	} else {
		return 8
	}
}
