package negativenumber

import "math"

func NegativeNumber(number int) int {
	return -int(math.Abs(float64(number)))
}
