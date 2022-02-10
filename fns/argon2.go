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
	usersalt := RANDOM_SALT_PLACEHOLDER
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(1)
	keylength := uint32(32)

	cmd := &cobra.Command{
		Use:   "argon2",
		Short: `Perform password key derivation using "argon2".`,
		Run: func(cmd *cobra.Command, args []string) {
			password, err := ioutil.ReadAll(os.Stdin)
			fatal(err, "failed to read password from stdin")

			var salt []byte
			if usersalt == RANDOM_SALT_PLACEHOLDER {
				salt = randombytes(16)
			} else {
				if salt, err = hex.DecodeString(usersalt); err != nil {
					fatal(err, "invalid salt hex encoding")
				}
			}

			bs := argon2.Key(password, salt, time, memory, threads, keylength)
			fmt.Println(hex.EncodeToString(bs))
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&usersalt, "salt", "s", usersalt, "salt value encoded as hexadecimal")
	flags.Uint32VarP(&time, "time", "t", time, "number of passes over the memory")
	flags.Uint32VarP(&memory, "memory", "m", memory, "size of the memory in KiB")
	flags.Uint8VarP(&threads, "threads", "j", threads, "number of concurrent threads to use")
	flags.Uint32VarP(&keylength, "length", "l", keylength, "length of resulting key in bytes")

	return cmd
}
