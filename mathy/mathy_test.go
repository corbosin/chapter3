package mathy

import "testing"

func TestMathy(t *testing.T) {
	var v float64
	v = Min([]float64{2, 5, 1, 6})
	if v != 1 {
		t.Error("Expected 1, got", v)
	}
}