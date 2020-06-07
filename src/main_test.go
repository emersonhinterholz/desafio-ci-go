package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	tests := []struct {
		x    int
		y    int
		want int
	}{
		{x: 1, y: 1, want: 2},
		{x: 2, y: 3, want: 5},
		{x: 2, y: 4, want: 6},
	}
	for _, tt := range tests {
		if got := sum(tt.x, tt.y); got != tt.want {
			t.Errorf("sum(%v, %v) = %v, want %v", tt.x, tt.y, got, tt.want)
		}
	}
}
