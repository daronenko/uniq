package cli_test

import (
	"reflect"
	"testing"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/args"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func TestGetModifiers(t *testing.T) {
	tests := []struct {
		name     string
		flags    args.Flags
		expected []modifier.Modifier
	}{
		{
			name: "No modifiers",
			flags: args.Flags{
				Count:      false,
				Repeated:   false,
				Unique:     false,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: []modifier.Modifier{},
		},
		{
			name: "IgnoreCase modifier only",
			flags: args.Flags{
				Count:      false,
				Repeated:   false,
				Unique:     false,
				IgnoreCase: true,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: []modifier.Modifier{modifier.NewIgnoreCaseModifier()},
		},
		{
			name: "SkipChars modifier only",
			flags: args.Flags{
				Count:      false,
				Repeated:   false,
				Unique:     false,
				IgnoreCase: false,
				SkipChars:  5,
				SkipFields: 0,
			},
			expected: []modifier.Modifier{modifier.NewSkipCharsModifier(5)},
		},
		{
			name: "SkipFields modifier only",
			flags: args.Flags{
				Count:      false,
				Repeated:   false,
				Unique:     false,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 2,
			},
			expected: []modifier.Modifier{modifier.NewSkipFieldsModifier(2)},
		},
		{
			name: "CountModifier only",
			flags: args.Flags{
				Count:      true,
				Repeated:   false,
				Unique:     false,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: []modifier.Modifier{modifier.NewCountModifier()},
		},
		{
			name: "Combination of modifiers",
			flags: args.Flags{
				Count:      true,
				Repeated:   false,
				Unique:     true,
				IgnoreCase: true,
				SkipChars:  3,
				SkipFields: 1,
			},
			expected: []modifier.Modifier{
				modifier.NewIgnoreCaseModifier(),
				modifier.NewSkipCharsModifier(3),
				modifier.NewSkipFieldsModifier(1),
				modifier.NewCountModifier(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.GetModifiers(&tt.flags)

			if len(got) != len(tt.expected) {
				t.Errorf("got %d modifiers, expected %d modifiers", len(got), len(tt.expected))
				return
			}

			for i := range got {
				if !reflect.DeepEqual(got[i], tt.expected[i]) {
					t.Errorf("got '%v', expected '%v'", got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestNewModifier(t *testing.T) {
	tests := []struct {
		name      string
		modifiers []modifier.Modifier
		expected  modifier.Modifier
	}{
		{
			name:      "No modifiers",
			modifiers: []modifier.Modifier{},
			expected:  modifier.NewMultiModifier([]modifier.Modifier{}),
		},
		{
			name: "Single modifier",
			modifiers: []modifier.Modifier{
				modifier.NewIgnoreCaseModifier(),
			},
			expected: modifier.NewMultiModifier([]modifier.Modifier{
				modifier.NewIgnoreCaseModifier(),
			}),
		},
		{
			name: "Multiple modifiers",
			modifiers: []modifier.Modifier{
				modifier.NewIgnoreCaseModifier(),
				modifier.NewSkipCharsModifier(5),
			},
			expected: modifier.NewMultiModifier([]modifier.Modifier{
				modifier.NewIgnoreCaseModifier(),
				modifier.NewSkipCharsModifier(5),
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.NewModifier(tt.modifiers)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got '%v', expected '%v'", got, tt.expected)
			}
		})
	}
}
