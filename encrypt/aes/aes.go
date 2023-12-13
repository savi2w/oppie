package aes

import (
	"crypto/aes"
	"io"
	"os"

	"github.com/savi2w/oppie/encrypt/kyber"
	"golang.org/x/crypto/xts"
)

const EncryptExtension = ".opp"
const BufferSize = 2 * 1024 * 1024 // 2MB

func File(loc string, k *kyber.Kyber) (esK string, err error) {
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

	cipher, err := xts.NewCipher(aes.NewCipher, sK)
	if err != nil {
		return "", err
	}

	pT := make([]byte, BufferSize)

	for {
		len, err := in.Read(pT)
		if err != nil && err != io.EOF {
			return "", err
		}

		if len == 0 {
			break
		}

		cT := make([]byte, len)

		cipher.Encrypt(cT, pT[:len], 0)

		if _, err := out.Write(cT); err != nil {
			return "", err
		}
	}

	return esK, nil
}
