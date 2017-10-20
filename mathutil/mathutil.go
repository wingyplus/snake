package mathutil

import "math"

// Constrain a value between a minimum and maximum value.
// NOTE: It port from p5.js see https://github.com/processing/p5.js/blob/master/src/math/calculation.js#L76
// for more detail.
func Constrain(n, low, high float64) float64 {
	return math.Max(math.Min(n, high), low)
}

func Dist(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x2-x1, y2-y1)
}
