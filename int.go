package collect

import "strconv"

func countInt(i int) int {
	str := strconv.Itoa(i)
	return len(str)
}
