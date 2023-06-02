package collect

import (
	"strings"
)

func countString(s string) int {
	var (
		res int
	)

	if s == "" {
		return 0
	}

	sp := strings.Split(s, "")
	for i := range sp {
		res = i + 1
	}
	return res
}
