package format

import (
	"fmt"

	"github.com/daronenko/uniq/pkg/uniq/modifier"
)

type Formatter interface {
	Format(line string) string
}

type DefaultFormat struct {
}

func NewDefaultFormat() *DefaultFormat {
	return &DefaultFormat{}
}

func (f *DefaultFormat) Format(line string) string {
	return line
}

type CountFormat struct {
	countModifier *modifier.CountModifier
}

func NewCountFormat(countModifier *modifier.CountModifier) *CountFormat {
	return &CountFormat{countModifier}
}

func (f *CountFormat) Format(line string) string {
	// function is called for the first line that
	// does not satisfy the uniqer condition
	defer f.countModifier.Set(1)
	return fmt.Sprintf("%d %s", f.countModifier.Count()-1, line)
}
