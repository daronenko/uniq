package uniqer_test

import (
	"testing"

	"github.com/daronenko/uniq/pkg/uniq/uniqer"
)

func TestDefaultUniqer(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected []bool
	}{
		{
			name:     "Different lines",
			lines:    []string{"line1", "line2", "line3"},
			expected: []bool{false, true, true},
		},
		{
			name:     "Repeated lines",
			lines:    []string{"line1", "line1", "line2", "line2"},
			expected: []bool{false, false, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := uniqer.NewDefaultUniqer()
			for i, line := range tt.lines {
				result := u.ShouldSave(line)
				if result != tt.expected[i] {
					t.Errorf("Expected %v, got %v for line %s", tt.expected[i], result, line)
				}
			}
		})
	}
}

func TestUniqueUniqer(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected []bool
	}{
		{
			name:     "Unique lines",
			lines:    []string{"line1", "line2", "line3"},
			expected: []bool{false, true, true, true},
		},
		{
			name:     "Repeated lines",
			lines:    []string{"line1", "line1", "line2", "line2"},
			expected: []bool{false, false, false, false, false},
		},
		{
			name:     "Unique and repeated mix",
			lines:    []string{"line1", "line1", "line2", "line3", "line3", "line4"},
			expected: []bool{false, false, false, true, false, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := uniqer.NewUniqueUniqer()
			for i, line := range tt.lines {
				result := u.ShouldSave(line)
				if result != tt.expected[i] {
					t.Errorf("Expected %v, got %v for line %s", tt.expected[i], result, line)
				}
			}

			if u.Finish() != tt.expected[len(tt.expected)-1] {
				t.Errorf("End() expected %v, got %v", tt.expected[len(tt.expected)-1], u.Finish())
			}
		})
	}
}

func TestRepeatedUniqer(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected []bool
	}{
		{
			name:     "Unique lines",
			lines:    []string{"line1", "line2", "line3"},
			expected: []bool{false, false, false, false},
		},
		{
			name:     "Repeated lines",
			lines:    []string{"line1", "line1", "line2", "line2", "line3"},
			expected: []bool{false, false, true, false, true, false},
		},
		{
			name:     "Unique and repeated mix",
			lines:    []string{"line1", "line1", "line2", "line3", "line3", "line4"},
			expected: []bool{false, false, true, false, false, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := uniqer.NewRepeatedUniqer()
			for i, line := range tt.lines {
				result := u.ShouldSave(line)
				if result != tt.expected[i] {
					t.Errorf("Expected %v, got %v for line %s", tt.expected[i], result, line)
				}
			}

			if u.Finish() != tt.expected[len(tt.expected)-1] {
				t.Errorf("End() expected %v, got %v", tt.expected[len(tt.expected)-1], u.Finish())
			}
		})
	}
}
