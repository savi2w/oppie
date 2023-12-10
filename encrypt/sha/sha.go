package sha

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const RandomSize = 128

func Hexadecimal(message []byte) (result string, err error) {
	hasher := sha256.New()
	hasher.Write(message)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func Random() (result []byte, err error) {
	random := make([]byte, RandomSize)

	if _, err := rand.Read(random); err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write(random)

	return hasher.Sum(nil), nil
}
