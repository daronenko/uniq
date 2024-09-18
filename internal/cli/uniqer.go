package cli

import (
	"github.com/daronenko/uniq/internal/cli/args"
	"github.com/daronenko/uniq/pkg/uniq/uniqer"
)

func NewUniqer(flags *args.Flags) uniqer.Uniqer {
	switch {
	case flags.Unique:
		return uniqer.NewUniqueUniqer()
	case flags.Repeated:
		return uniqer.NewRepeatedUniqer()
	default:
		return uniqer.NewDefaultUniqer()
	}
}
