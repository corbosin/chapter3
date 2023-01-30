package main

import "fmt"
import "./mathy"

func main() {
	i := []float64{64, 5, 67, 24, 75,
	}

	minimal := mathy.Min(i)
	maximum := mathy.Max(i)
	fmt.Println("Minimal figure is - ", minimal)
	fmt.Println("Maximum figure is - ", maximum)
}