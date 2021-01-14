package main

import (
	"testing"
)

func TestXplusY(t *testing.T) {
	tcases := map[string]struct {
		x    float32
		y    float32
		want float32
	}{
		"t1": {2.5, 1.5, 4.0},
		"t2": {3.1, 1.2, 4.3},
	}
	g, _ := gorXplusYGraph()

	for name, tc := range tcases {
		t.Run(name, func(t *testing.T) {
			got, _ := gorXplusY(g, tc.x, tc.y)
			if tc.want != got {
				t.Errorf("want: %v got: %v", tc.want, got)
			}
		})

	}
}
