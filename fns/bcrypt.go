package fns

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func GetBcryptCmd() *cobra.Command {
	cost := bcrypt.DefaultCost
	var verify string

	cmd := &cobra.Command{
		Use:   "bcrypt",
		Short: `Perform password key derivation and verification using "bcrypt".`,
		Run: func(cmd *cobra.Command, args []string) {
			password := readUserPassword()

			if verify == "" {
				bs, err := bcrypt.GenerateFromPassword(password, cost)
				fatal(err, "failed to generate bcrypt hash")
				fmt.Println(string(bs))
			} else {
				err := bcrypt.CompareHashAndPassword([]byte(verify), password)
				if err != nil {
					log.Fatal("ERROR: verify string is not the hash of the given password")
					os.Exit(1)
				}
				log.Printf("OK: verify string matches given password")
				os.Exit(0)
			}
		},
	}

	flags := cmd.Flags()
	flags.IntVarP(&cost, "cost", "c", cost,
		fmt.Sprintf("cost parameter in [%v,%v]", bcrypt.MinCost, bcrypt.MaxCost))
	flags.StringVarP(&verify, "verify", "v", verify,
		"bcrypt hashed password for comparison with possible plaintext equivalent")

	return cmd
}
