package cli

import (
	"github.com/daronenko/uniq/internal/cli/format"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func NewFormat(modifiers []modifier.Modifier) format.Formatter {
	for _, f := range modifiers {
		switch v := f.(type) {
		case *modifier.CountModifier:
			return format.NewCountFormat(v)
		}
	}

	return format.NewDefaultFormat()
}
