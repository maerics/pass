package main

import (
	"bytes"
	"regexp"
	"testing"
)

func TestMainHelp(t *testing.T) {
	helpArgs := [][]string{{}, {"-h"}, {"--help"}, {"help"}}

	helpMessages := []*regexp.Regexp{
		regexp.MustCompile(`^Password hashing and key derivation utilities.\n`),
		regexp.MustCompile(`Usage:\n  pass \[command\]`),
		regexp.MustCompile(`Flags:\n  -`),
		regexp.MustCompile(`\nUse "pass \[command\] --help" for more information about a command.\n$`),
	}

	for _, args := range helpArgs {
		buf := &bytes.Buffer{}
		passCmd = newPassCmd()
		passCmd.SetArgs(args)
		passCmd.SetIn(nil)
		passCmd.SetOut(buf)
		main()
		output := buf.String()

		for _, helpMessage := range helpMessages {
			if !helpMessage.MatchString(output) {
				t.Fatalf(
					"unexpected help message\n:  wanted: %q\n     got: %q",
					helpMessage.String(), output)
			}
		}
	}
}
