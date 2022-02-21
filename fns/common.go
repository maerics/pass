package fns

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/term"
)

const RANDOM_SALT_PLACEHOLDER = "random"

func readUserPassword() []byte {
	stdinfd := int(os.Stdin.Fd())

	if term.IsTerminal(stdinfd) {
		// Read from terminal without echo
		fmt.Fprint(os.Stderr, "Password: ")
		password, err := term.ReadPassword(stdinfd)
		fatal(err, "failed to read password from user input")
		fmt.Fprintln(os.Stderr)
		return password
	}

	// Read per usual.
	password, err := ioutil.ReadAll(os.Stdin)
	fatal(err, "failed to read password from stdin")
	return password
}

func randombytes(n int) []byte {
	nonce := make([]byte, n)
	if m, err := rand.Read(nonce); err != nil || m != len(nonce) {
		fatal(err, "failed to read random bytes")
	}
	return nonce
}

func getSalt(hexbytes string) []byte {
	var err error
	var salt []byte
	if hexbytes == RANDOM_SALT_PLACEHOLDER {
		salt = randombytes(16)
	} else {
		if salt, err = hex.DecodeString(hexbytes); err != nil {
			fatal(err, "invalid salt hex encoding")
		}
	}
	return salt
}

func fatal(err error, message string) {
	if err != nil {
		log.Fatalf("FATAL: %v: %v", message, err)
	}
}
