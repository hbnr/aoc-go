package main

import "testing"

func TestValidId(t *testing.T) {
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
	}

	for _, tt := range tests {
		result := validId(tt.id)
		if result != tt.expected {
			t.Errorf("ValidId(%v) = %v; want %v", tt.id, result, tt.expected)
		}
	}
}
