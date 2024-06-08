package utils

import "crypto/sha256"

func SHA256Hash(d []byte) []byte {
	hash := sha256.Sum256(d)
	return hash[:]
}
