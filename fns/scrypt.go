package fns

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/scrypt"
)

func GetScryptCmd() *cobra.Command {
	usersalt := RANDOM_SALT_PLACEHOLDER
	n, r, p := 32768, 8, 1
	keylength := 32

	cmd := &cobra.Command{
		Use:   "scrypt",
		Short: `Perform password key derivation using "scrypt".`,
		Run: func(cmd *cobra.Command, args []string) {
			password := readUserPassword()
			salt := getSalt(usersalt)
			bs, err := scrypt.Key(password, salt, n, r, p, keylength)
			fatal(err, "scrypt key derivation failed")
			fmt.Println(hex.EncodeToString(bs))
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&usersalt, "salt", "s", usersalt,
		"salt value encoded as hexadecimal")
	flags.IntVarP(&n, "N", "N", n,
		"CPU/memory cost parameter, which must be a power of two greater than 1")
	flags.IntVarP(&r, "r", "r", r,
		"r and p must satisfy r * p < 2^30")
	flags.IntVarP(&p, "p", "p", p,
		"r and p must satisfy r * p < 2^30")
	flags.IntVarP(&keylength, "length", "l", keylength,
		"length of resulting key in bytes")

	return cmd
}
