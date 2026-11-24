package main

import "testing"

func TestIsValidTree(t *testing.T) {
	cases := []struct {
		n     int
		edges [][2]int
		want  bool
	}{
		{5, [][2]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{5, [][2]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false}, // cycle
		{4, [][2]int{{0, 1}, {2, 3}}, false},                         // right edge count needs n-1: 2 != 3
		{4, [][2]int{{0, 1}, {2, 3}, {0, 1}}, false},                 // n-1 edges but disconnected (duplicate edge)
		{1, nil, true}, // single node, no edges
		{2, [][2]int{{0, 1}}, true},
		{0, nil, false},
	}
	for _, c := range cases {
		if got := IsValidTree(c.n, c.edges); got != c.want {
			t.Errorf("IsValidTree(%d,%v) = %v, want %v", c.n, c.edges, got, c.want)
		}
	}
}
