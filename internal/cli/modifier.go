package cli

import (
	"github.com/daronenko/uniq/internal/cli/args"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

func GetModifiers(flags *args.Flags) []modifier.Modifier {
	modifiers := []modifier.Modifier{}

	if flags.IgnoreCase {
		modifiers = append(modifiers, modifier.NewIgnoreCaseModifier())
	}

	if flags.SkipChars > 0 {
		modifiers = append(modifiers, modifier.NewSkipCharsModifier(flags.SkipChars))
	}

	if flags.SkipFields > 0 {
		modifiers = append(modifiers, modifier.NewSkipFieldsModifier(flags.SkipFields))
	}

	if flags.Count {
		modifiers = append(modifiers, modifier.NewCountModifier())
	}

	return modifiers
}

func NewModifier(modifiers []modifier.Modifier) modifier.Modifier {
	return modifier.NewMultiModifier(modifiers)
}
