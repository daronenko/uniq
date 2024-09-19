package args

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type IOStream struct {
	Input  io.Reader
	Output io.Writer
}

func NewIOStream() (*IOStream, error) {
	arguments := &IOStream{}

	err := arguments.parse()
	if err != nil {
		return nil, err
	}

	return arguments, nil
}

func (ios *IOStream) Close() error {
	if err := closeImpl(ios.Input); err != nil {
		return fmt.Errorf("cannot close input stream: %v", err)
	}

	if err := closeImpl(ios.Output); err != nil {
		return fmt.Errorf("cannot close output stream: %v", err)
	}

	return nil
}

func closeImpl(stream interface{}) error {
	if _, ok := stream.(io.Closer); !ok {
		return fmt.Errorf("no way to close the stream")
	}

	if err := stream.(io.Closer).Close(); err != nil {
		return err
	}

	return nil
}

func (ios *IOStream) parse() error {
	arguments := flag.Args()

	var inputStream, outputStream *os.File
	var err error

	if len(arguments) > 0 {
		inputStream, err = os.Open(arguments[0])
		if err != nil {
			return fmt.Errorf("cannot setup input stream: %v", err)
		}
	} else {
		inputStream = os.Stdin
	}

	if len(arguments) > 1 {
		outputStream, err = os.OpenFile(arguments[1], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("cannot setup output stream: %v", err)
		}
	} else {
		outputStream = os.Stdout
	}

	ios.Input = inputStream
	ios.Output = outputStream

	return nil
}
