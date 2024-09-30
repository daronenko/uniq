package cli

import (
	"github.com/daronenko/uniq/pkg/uniq"
	"github.com/daronenko/uniq/pkg/uniq/modifier"
	"github.com/daronenko/uniq/pkg/uniq/uniqer"
)

func NewUniq(uniqer uniqer.Uniqer, modifier modifier.Modifier) *uniq.Uniq {
	return uniq.New(uniqer, modifier)
}
