package input

import "testing"

func TestGetFileName(t *testing.T) {
	tests := []struct {
		sample          bool
		year, day, want string
	}{
		{true, "2025", "7", "cmd/2025/day07/sample.txt"},
		{false, "2025", "7", "cmd/2025/day07/input.txt"},
		{true, "2025", "10", "cmd/2025/day10/sample.txt"},
		{false, "2025", "10", "cmd/2025/day10/input.txt"},
	}

	for _, tt := range tests {
		got := GetFileName(tt.year, tt.day, tt.sample)
		if got != tt.want {
			t.Errorf("GetFileName(%v, %v, %v) = %v, want %v", tt.year, tt.day, tt.sample, got, tt.want)
		}
	}
}
