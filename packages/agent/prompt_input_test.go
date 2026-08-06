package agent

import "testing"

func TestCombinePromptInput(t *testing.T) {
	tests := []struct {
		name       string
		stdin      string
		positional string
		want       string
	}{
		{name: "stdin and instruction", stdin: "input text\n", positional: "summarize it", want: "input text\nsummarize it"},
		{name: "stdin only", stdin: " input text\n", want: "input text"},
		{name: "positional only", positional: " summarize it ", want: "summarize it"},
		{name: "empty", stdin: " \n", positional: " ", want: ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := combinePromptInput(tc.stdin, tc.positional); got != tc.want {
				t.Fatalf("combinePromptInput(%q, %q) = %q, want %q", tc.stdin, tc.positional, got, tc.want)
			}
		})
	}
}
