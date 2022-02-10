package fns

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

const RANDOM_SALT_PLACEHOLDER = "random"

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
