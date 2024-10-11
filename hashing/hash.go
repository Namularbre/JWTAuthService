package hashing

import (
	"bytes"
	"golang.org/x/crypto/argon2"
	"os"
)

func Hash(plain string) []byte {
	salt := []byte(os.Getenv("SALT"))
	return argon2.IDKey([]byte(plain), salt, 1, 64*1024, 4, 32)
}

func Compare(plain string, hash []byte) bool {
	return bytes.Compare(Hash(plain), hash) == 0
}
