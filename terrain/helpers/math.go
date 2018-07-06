package helpers

func Constrain(num float64) float64 {
	if num < 0 {
		return 0.0
	}
	if num > 1 {
		return 1.0
	}
	return num
}
