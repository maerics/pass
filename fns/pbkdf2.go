package fns

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/pbkdf2"
)

func GetPbkdf2Cmd() *cobra.Command {
	usersalt := RANDOM_SALT_PLACEHOLDER
	iterations := 4096
	keylength := 32
	hash := "sha256"

	cmd := &cobra.Command{
		Use:   "pbkdf2",
		Short: `Perform password key derivation using "pbkdf2".`,
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

			hashfn := getPbkdf2HashFn(hash)
			bs := pbkdf2.Key(password, salt, iterations, keylength, hashfn)
			fmt.Println(hex.EncodeToString(bs))
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&usersalt, "salt", "s", usersalt, "salt value encoded as hexadecimal")
	flags.IntVarP(&iterations, "iterations", "i", iterations, "number of iterations to perform")
	flags.IntVarP(&keylength, "length", "l", keylength, "length of resulting key in bytes")
	flags.StringVarP(&hash, "hash", "H", hash, "hash function for HMAC computatiton")

	return cmd
}

func getPbkdf2HashFn(name string) func() hash.Hash {
	n := strings.ToLower(regexp.MustCompile(`\W`).ReplaceAllString(name, ""))
	if hash, present := pbkdf2Hashes[n]; present {
		return hash
	}
	log.Fatalf("unknown hash function %q", name)
	return nil
}

var pbkdf2Hashes = map[string]func() hash.Hash{
	"sha1":   sha1.New,
	"sha224": sha512.New512_224,
	"sha256": sha256.New,
	"sha384": sha512.New384,
	"sha512": sha512.New,
}
