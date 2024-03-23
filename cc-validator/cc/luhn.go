package cc

import (
	"strconv"
	"strings"
)

func isValidLuhn(c string) bool {
	digits := strings.Split(c, "")
	var checkDigit int
	var sum int

	for i := len(digits) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(digits[i], 10, 8)

		if err != nil {
			return false
		}

		if i == len(digits)-1 {
			checkDigit = int(n)
			continue
		}

		if i%2 == 0 {
			n = n * 2
			if n > 9 {
				n = n - 9
			}
		}
		sum += int(n)
	}

	return (sum+checkDigit)%10 == 0
}
