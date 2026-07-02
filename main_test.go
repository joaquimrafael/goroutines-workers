package main

import "testing"

func TestRodarPool(t *testing.T) {
	in := []int{2, 3, 1, 12, 77, 10, 131, 120}
	out := RunPool(in, 3)

	if len(out) != len(in) {
		t.Errorf("esperava %d resultados, veio %d", len(in), len(out))
	}
}
