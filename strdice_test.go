package strdice

import (
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   float64
	}{
		{"", "", 1},
		{" ", " ", 1},
		{"a", "b", 0},
		{"a", "a", 1},
		{"a", "", 0},
		{" ", "a", 0},
		{"bad apple", "badapple", 1},
		{"this is a string", "this is not a string", 0.8148148148148148},
		{"is that a cat?", "is that a dog?", 0.6},
		{"larder", "lerder", 0.6},
		{"mouse", "mouse trap", 0.6666666666666666},
		{"mouse trap", "mouse", 0.6666666666666666},
		{"i want something to play today", "i want something to draw tomorrow", 0.6666666666666666},
		{"Just because the water is red doesn't mean you can't drink it.", "Just because the water is blue doesn't mean you can't drink it.", 0.9306930693069307},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s,%s", test.s1, test.s2)
		t.Run(name, func(t *testing.T) {
			got := Compute(test.s1, test.s2)

			if got != test.want {
				t.Errorf("got %f, wanted %f", got, test.want)
			}
		})
	}
}
