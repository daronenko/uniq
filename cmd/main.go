package main

import (
	"fmt"
	"os"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/args"
)

func main() {
	flags, err := args.NewFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "flags error: %v\n", err)
		os.Exit(1)
	}

	iostream, err := args.NewIOStream()
	if err != nil {
		fmt.Fprintf(os.Stderr, "iostream error: %v\n", err)
		os.Exit(1)
	}

	cmd := cli.New(flags, iostream)
	cmd.Run()
}
