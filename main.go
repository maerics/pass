package main

import (
	"log"
	"os"
	"pass/fns"

	"github.com/spf13/cobra"
)

func newPassCmd() *cobra.Command {
	return &cobra.Command{
		Use: "pass", Short: "Password hashing and key derivation utilities.",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
}

var passCmd = newPassCmd()

func main() {
	log.SetFlags(0)

	for _, cmd := range []*cobra.Command{
		fns.GetArgon2Cmd(),
		fns.GetPbkdf2Cmd(),
		fns.GetBcryptCmd(),
		fns.GetScryptCmd(),
	} {
		passCmd.AddCommand(cmd)
	}

	if err := passCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
