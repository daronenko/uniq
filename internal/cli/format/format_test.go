package format_test

import (
	"testing"

	"github.com/daronenko/uniq/internal/cli/format"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func TestDefaultFormat(t *testing.T) {
	formatter := format.NewDefaultFormat()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Simple string",
			input:    "Hello",
			expected: "Hello",
		},
		{
			name:     "String with spaces",
			input:    "Hello World",
			expected: "Hello World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatter.Format(tt.input)
			if got != tt.expected {
				t.Errorf("got '%v', expected '%v'", got, tt.expected)
			}
		})
	}
}

func TestCountFormat(t *testing.T) {
	countModifier := modifier.NewCountModifier()
	formatter := format.NewCountFormat(countModifier)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "First call",
			input:    "Hello",
			expected: "0 Hello",
		},
		{
			name:     "Second call",
			input:    "World",
			expected: "1 World",
		},
		{
			name:  "Third call",
			input: "World",
			// format resets the counter because it is called once
			// when the sequence finishes
			expected: "1 World",
		},
		{
			name:     "Fourth call",
			input:    "Foo",
			expected: "1 Foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countModifier.Modify(tt.input)
			got := formatter.Format(tt.input)
			if got != tt.expected {
				t.Errorf("got '%v', expected '%v'", got, tt.expected)
			}
		})
	}
}
