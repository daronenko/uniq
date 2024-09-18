package modifier_test

import (
	"testing"

	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func TestCountModifier(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected uint
	}{
		{
			name:     "Three lines",
			lines:    []string{"line1", "line2", "line3"},
			expected: 3,
		},
		{
			name:     "No lines",
			lines:    []string{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := modifier.NewCountModifier()

			for _, line := range tt.lines {
				m.Modify(line)
			}

			if m.Count() != tt.expected {
				t.Errorf("got %d, expected %d", m.Count(), tt.expected)
			}
		})
	}
}

func TestSkipFieldsModifier(t *testing.T) {
	tests := []struct {
		name        string
		fieldsCount uint
		line        string
		expected    string
	}{
		{
			name:        "Skip 2 fields",
			fieldsCount: 2,
			line:        "one two three four",
			expected:    "three four",
		},
		{
			name:        "Less fields than needed",
			fieldsCount: 2,
			line:        "one",
			expected:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := modifier.NewSkipFieldsModifier(tt.fieldsCount)
			got := m.Modify(tt.line)
			if got != tt.expected {
				t.Errorf("got '%s', expected '%s'", got, tt.expected)
			}
		})
	}
}

func TestSkipCharsModifier(t *testing.T) {
	tests := []struct {
		name       string
		charsCount uint
		line       string
		expected   string
	}{
		{
			name:       "Skip 5 chars",
			charsCount: 5,
			line:       "123456789",
			expected:   "6789",
		},
		{
			name:       "Fewer chars than needed",
			charsCount: 5,
			line:       "1234",
			expected:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := modifier.NewSkipCharsModifier(tt.charsCount)
			got := m.Modify(tt.line)
			if got != tt.expected {
				t.Errorf("got '%s', expected '%s'", got, tt.expected)
			}
		})
	}
}

func TestIgnoreCaseModifier(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{
			name:     "Mixed case string",
			line:     "HeLLo WoRLd",
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := modifier.NewIgnoreCaseModifier()
			got := m.Modify(tt.line)
			if got != tt.expected {
				t.Errorf("got '%s', expected '%s'", got, tt.expected)
			}
		})
	}
}

func TestMultipleModifier(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
	}{
		{
			name:     "SkipFields and IgnoreCase modifiers",
			line:     "HeLLo WoRLd",
			expected: "world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skipFieldsModifier := modifier.NewSkipFieldsModifier(1)
			ignoreCaseModifier := modifier.NewIgnoreCaseModifier()
			multiModifier := modifier.NewMultiModifier([]modifier.Modifier{skipFieldsModifier, ignoreCaseModifier})

			got := multiModifier.Modify(tt.line)
			if got != tt.expected {
				t.Errorf("got '%s', expected '%s'", got, tt.expected)
			}
		})
	}
}
