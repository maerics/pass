package fns

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/scrypt"
)

func GetScryptCmd() *cobra.Command {
	options := struct {
		Salt      string
		N, r, p   int
		KeyLength int
	}{
		Salt: RANDOM_SALT_PLACEHOLDER, N: 32768, r: 8, p: 1, KeyLength: 32,
	}

	cmd := &cobra.Command{
		Use:   "scrypt",
		Short: `Perform password key derivation using "scrypt".`,
		Run: func(cmd *cobra.Command, args []string) {
			password, err := ioutil.ReadAll(os.Stdin)
			fatal(err, "failed to read password from stdin")

			var salt []byte
			if options.Salt == RANDOM_SALT_PLACEHOLDER {
				salt = randombytes(16)
			} else {
				if salt, err = hex.DecodeString(options.Salt); err != nil {
					fatal(err, "invalid salt hex encoding")
				}
			}

			bs, err := scrypt.Key(password, salt, options.N, options.r, options.p, options.KeyLength)
			fatal(err, "scrypt key generation failed")
			fmt.Println(hex.EncodeToString(bs))
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.Salt, "salt", "s", options.Salt,
		"salt value encoded as hexadecimal")
	flags.IntVarP(&options.N, "N", "N", options.N,
		"CPU/memory cost parameter, which must be a power of two greater than 1")
	flags.IntVarP(&options.r, "r", "r", options.r,
		"r and p must satisfy r * p < 2^30")
	flags.IntVarP(&options.p, "p", "p", options.p,
		"r and p must satisfy r * p < 2^30")
	flags.IntVarP(&options.KeyLength, "length", "l", options.KeyLength,
		"length of resulting key in bytes")

	return cmd
}
