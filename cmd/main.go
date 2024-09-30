package main

import (
	"log"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/args"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	flags, err := args.NewFlags()
	if err != nil {
		log.Fatalf("flags error: %v\n", err)
	}

	iostream, err := args.NewIOStream()
	if err != nil {
		log.Fatalf("iostream error: %v\n", err)
	}
	defer iostream.Close()

	cmd := cli.New(flags, iostream)
	cmd.Run()
}
