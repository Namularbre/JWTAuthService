package hashing

import (
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"os"
	"strings"
)

func Hash(plain string) string {
	salt := []byte(os.Getenv("SALT"))
	return hex.EncodeToString(argon2.IDKey([]byte(plain), salt, 1, 64*1024, 4, 32))
}

func Compare(plain string, hash string) bool {
	return strings.Compare(Hash(plain), hash) == 0
}
