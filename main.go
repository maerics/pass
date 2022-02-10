package main

import (
	"log"
	"os"
	"pass/fns"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)

	rootCmd := &cobra.Command{
		Use: "pass", Short: "Password hashing and key derivation utilities.",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	for _, cmd := range []*cobra.Command{
		fns.GetArgon2Cmd(),
		fns.GetPbkdf2Cmd(),
		fns.GetBcryptCmd(),
		fns.GetScryptCmd(),
	} {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
