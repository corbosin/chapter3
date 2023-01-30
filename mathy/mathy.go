package mathy

func Min(x []float64) float64 {
	
min := x[0]
	for _, y := range x {
		if min > y {
			min = y
		}
	}
	return min
}


func Max(x []float64) float64 {
	
	max := x[0]
		for _, y := range x {
			if max < y {
				max = y
			}
		}
		return max
	}