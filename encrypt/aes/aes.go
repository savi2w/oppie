package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	"github.com/savi2w/oppie/encrypt/kyber"
)

const EncryptExtension = ".opp"
const BufferSize = 2 * 1024 * 1024 // 2MB

func Counter(loc string, k *kyber.Kyber) (esK string, err error) {
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

	sK, esK, err := k.SecretKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(sK)
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

	pT := make([]byte, BufferSize)
	stream := cipher.NewCTR(block, nonce)

	for {
		len, err := in.Read(pT)
		if err != nil && err != io.EOF {
			return "", err
		}

		if len == 0 {
			break
		}

		cT := make([]byte, len)

		stream.XORKeyStream(cT, pT[:len])

		if _, err := out.Write(cT); err != nil {
			return "", err
		}
	}

	if _, err := io.ReadFull(rand.Reader, sK); err != nil {
		return "", err
	}

	return esK, nil
}
