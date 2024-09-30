package cli_test

import (
	"reflect"
	"testing"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/format"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func TestNewFormat(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []modifier.Modifier
		expected  format.Formatter
	}{
		{
			name:      "CountModifier only",
			modifiers: []modifier.Modifier{modifier.NewCountModifier()},
			expected:  format.NewCountFormat(modifier.NewCountModifier()),
		},
		{
			name:      "Multiple modifiers with CountModifier",
			modifiers: []modifier.Modifier{modifier.NewSkipCharsModifier(5), modifier.NewCountModifier()},
			expected:  format.NewCountFormat(modifier.NewCountModifier()),
		},
		{
			name:      "No CountModifier",
			modifiers: []modifier.Modifier{modifier.NewSkipCharsModifier(5)},
			expected:  format.NewDefaultFormat(),
		},
		{
			name:      "Empty modifiers",
			modifiers: []modifier.Modifier{},
			expected:  format.NewDefaultFormat(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.NewFormat(tt.modifiers)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got '%v', expected '%v'", got, tt.expected)
			}
		})
	}
}
