package args

import (
	"flag"
	"fmt"
)

type Flags struct {
	Count      bool
	Repeated   bool
	Unique     bool
	IgnoreCase bool
	SkipChars  uint
	SkipFields uint
}

func NewFlags() (*Flags, error) {
	flags := &Flags{}

	flags.register()
	flags.parse()

	err := flags.validate()
	if err != nil {
		return nil, fmt.Errorf("invalid flags combination: %v", err)
	}

	return flags, nil
}

func (f *Flags) register() {
	flag.BoolVar(&f.Count, "c", false, "prefix lines by the number of occurrences")
	flag.BoolVar(&f.Repeated, "d", false, "only print duplicate lines, one for each group")
	flag.BoolVar(&f.Unique, "u", false, "only print unique lines")
	flag.BoolVar(&f.IgnoreCase, "i", false, "ignore differences in case when comparing")
	flag.UintVar(&f.SkipChars, "s", 0, "avoid comparing the first N characters")
	flag.UintVar(&f.SkipFields, "f", 0, "avoid comparing the first N fields")
}

func (f *Flags) parse() {
	flag.Parse()
}

func (f *Flags) validate() error {
	if (f.Count && f.Repeated) || (f.Count && f.Unique) || (f.Repeated && f.Unique) {
		return fmt.Errorf("only one of -c, -d, -u flags can be active")
	}

	return nil
}
