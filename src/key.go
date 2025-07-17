package main

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(passphrase, salt string) []byte {
	const keyLen = 32
	const iterations = 100_000
	key := pbkdf2.Key([]byte(passphrase), []byte(salt), iterations, keyLen, sha256.New)
	return key
}
