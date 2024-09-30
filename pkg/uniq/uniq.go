package uniq

import (
	"github.com/daronenko/uniq/pkg/uniq/modifier"
	"github.com/daronenko/uniq/pkg/uniq/uniqer"
)

type Uniq struct {
	uniqer   uniqer.Uniqer
	modifier modifier.Modifier
}

func New(uniqer uniqer.Uniqer, modifier modifier.Modifier) *Uniq {
	return &Uniq{
		uniqer,
		modifier,
	}
}

func (u *Uniq) Modify(line string) string {
	return u.modifier.Modify(line)
}

func (u *Uniq) ShouldSave(line string) bool {
	return u.uniqer.ShouldSave(line)
}

func (u *Uniq) Finish() bool {
	return u.uniqer.Finish()
}
