package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1HashHex(plain string) string {
	if plain == "" {
		return plain
	}

	h := sha1.New()
	h.Write([]byte(plain))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA1HashEquals(plain string, hash string) bool {
	return SHA1HashHex(plain) == hash
}
