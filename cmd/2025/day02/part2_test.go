package main

import "testing"

func TestValidId2(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{1, true},
		{11, false},
		{12, true},
		{1212, false},
		{1213, true},
		{1221, true},
		{99, false},
		{111, false},
		{222222, false},
		{2121212121, false},
	}

	for _, tt := range tests {
		result := validId2(tt.id)
		if result != tt.expected {
			t.Errorf("ValidId(%v) = %v; want %v", tt.id, result, tt.expected)
		}
	}
}
