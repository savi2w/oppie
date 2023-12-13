package sha

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func Hexadecimal(message []byte) (result string, err error) {
	bytes := sha3.Sum256(message)

	return hex.EncodeToString(bytes[:]), nil
}
