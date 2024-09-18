package cli_test

import (
	"reflect"
	"testing"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/args"
	"github.com/daronenko/uniq/pkg/uniq/uniqer"
)

func TestNewUniqer(t *testing.T) {
	tests := []struct {
		name     string
		flags    args.Flags
		expected uniqer.Uniqer
	}{
		{
			name: "Unique flag set",
			flags: args.Flags{
				Count:      false,
				Repeated:   false,
				Unique:     true,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: uniqer.NewUniqueUniqer(),
		},
		{
			name: "Repeated flag set",
			flags: args.Flags{
				Count:      false,
				Repeated:   true,
				Unique:     false,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: uniqer.NewRepeatedUniqer(),
		},
		{
			name:     "No flags set",
			flags:    args.Flags{},
			expected: uniqer.NewDefaultUniqer(),
		},
		{
			name: "Multiple flags, Unique wins",
			flags: args.Flags{
				Count:      false,
				Repeated:   true,
				Unique:     true,
				IgnoreCase: false,
				SkipChars:  0,
				SkipFields: 0,
			},
			expected: uniqer.NewUniqueUniqer(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.NewUniqer(&tt.flags)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got '%v', expected '%v'", got, tt.expected)
			}
		})
	}
}
