package main

import (
	"testing"
)

func TestXplusY(t *testing.T) {
	tcases := map[string]struct {
		x    float64
		y    float64
		want float64
	}{
		"t1": {2.5, 1.5, 4.0},
	}

	for name, tc := range tcases {
		t.Run(name, func(t *testing.T) {
			got, _ := gorXplusY(tc.x, tc.y)
			if tc.want != got {
				t.Errorf("want: %v got: %v", tc.want, got)
			}
		})

	}
}
