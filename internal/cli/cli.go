package cli

import (
	"bufio"
	"io"

	"github.com/daronenko/uniq/internal/cli/args"
	"github.com/daronenko/uniq/internal/cli/format"
	"github.com/daronenko/uniq/pkg/uniq"
)

type UniqCmd struct {
	*uniq.Uniq
	*args.IOStream
	formatter format.Formatter
}

func New(flags *args.Flags, iostream *args.IOStream) *UniqCmd {
	uniqer := NewUniqer(flags)

	modifiers := GetModifiers(flags)
	modifier := NewModifier(modifiers)

	uniq := NewUniq(uniqer, modifier)
	formatter := NewFormat(modifiers)

	return &UniqCmd{
		uniq,
		iostream,
		formatter,
	}
}

func (cmd *UniqCmd) Run() {
	// `target` is the beginning of the sequence we are searching for
	var targetLine, modifiedTargetLine string
	var notFirstLine bool

	scanner := bufio.NewScanner(cmd.Input)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := cmd.Modify(line)

		if !notFirstLine {
			targetLine = line
			modifiedTargetLine = modifiedLine
			notFirstLine = true
		}

		if cmd.ShouldSave(modifiedLine) {
			io.WriteString(cmd.Output, cmd.formatter.Format(targetLine)+"\n")
		}

		if modifiedLine != modifiedTargetLine {
			targetLine = line
		}
	}

	if cmd.Finish() {
		cmd.Modify(targetLine)
		io.WriteString(cmd.Output, cmd.formatter.Format(targetLine)+"\n")
	}
}
