package class_03_two_cristal_ball_problem

import "math"

func two_cristal_ball_problem(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	// 1. Use the first ball for find where it breaks
	i := jumpAmount
out:
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break out
		}
	}

	// 2.
	i -= jumpAmount
	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		i++
		if breaks[i] {
			return i
		}
	}
	return -1
}
