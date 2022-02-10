package fns

import (
	"crypto/rand"
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

func fatal(err error, message string) {
	if err != nil {
		log.Fatalf("FATAL: %v: %v", message, err)
	}
}

