package cli_test

import (
	"bytes"
	"testing"

	"github.com/daronenko/uniq/internal/cli"
	"github.com/daronenko/uniq/internal/cli/args"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		flags    args.Flags
	}{
		{
			name:  "without flags",
			flags: args.Flags{},
			input: `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`,
			expected: `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
`,
		},
		{
			name:  "-c",
			flags: args.Flags{Count: true},
			input: `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`,
			expected: `3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.
`,
		},
		{
			name:  "-d",
			flags: args.Flags{Repeated: true},
			input: `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`,
			expected: `I love music.
I love music of Kartik.
I love music of Kartik.
`,
		},
		{
			name:  "-u",
			flags: args.Flags{Unique: true},
			input: `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`,
			expected: `
Thanks.
`,
		},
		{
			name:  "-i",
			flags: args.Flags{IgnoreCase: true},
			input: `I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.
`,
			expected: `I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.
`,
		},
		{
			name:  "-f 1",
			flags: args.Flags{SkipFields: 1},
			input: `We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`,
			expected: `We love music.

I love music of Kartik.
Thanks.
`,
		},
		{
			name:  "-s 1",
			flags: args.Flags{SkipChars: 1},
			input: `I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`,
			expected: `I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBuffer := bytes.NewBufferString(tt.input)
			outputBuffer := new(bytes.Buffer)

			iostream := args.IOStream{
				Input:  inputBuffer,
				Output: outputBuffer,
			}

			cmd := cli.New(&tt.flags, &iostream)

			cmd.Run()

			if outputBuffer.String() != tt.expected {
				t.Errorf("\n--- expected:\n%s\n--- got:\n%s\n", tt.expected, outputBuffer.String())
			}
		})
	}
}
