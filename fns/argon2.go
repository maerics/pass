package fns

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/argon2"
)

func GetArgon2Cmd() *cobra.Command {
	options := struct {
		Salt         string
		Time, Memory uint32
		Threads      uint8
		KeyLength    uint32
	}{
		Salt: RANDOM_SALT_PLACEHOLDER,
		Time: 1, Memory: 64 * 1024, Threads: 1, KeyLength: 32,
	}

	cmd := &cobra.Command{
		Use:   "argon2",
		Short: `Perform password key derivation using "argon2".`,
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

			bs := argon2.Key(password, salt, options.Time, options.Memory, options.Threads, options.KeyLength)
			fmt.Println(hex.EncodeToString(bs))
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&options.Salt, "salt", "s", options.Salt,
		"salt value encoded as hexadecimal")
	flags.Uint32VarP(&options.Time, "time", "t", options.Time,
		"number of passes over the memory")
	flags.Uint32VarP(&options.Memory, "memory", "m", options.Memory,
		"size of the memory in KiB")
	flags.Uint8VarP(&options.Threads, "threads", "j", options.Threads,
		"number of concurrent threads to use")
	flags.Uint32VarP(&options.KeyLength, "length", "l", options.KeyLength,
		"length of resulting key in bytes")

	return cmd
}
