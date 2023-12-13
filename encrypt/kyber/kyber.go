package kyber

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"

	pke "github.com/cloudflare/circl/pke/kyber/kyber1024"
	"github.com/savi2w/oppie/config"
	"golang.org/x/crypto/sha3"
)

type Kyber struct {
	publicKey *pke.PublicKey
}

func New() (result *Kyber, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("[Kyber] Error unpacking public key")
		}
	}()

	kyber := &Kyber{
		publicKey: &pke.PublicKey{},
	}

	block, _ := pem.Decode([]byte(config.PublicKey))
	kyber.publicKey.Unpack(block.Bytes)

	return kyber, nil
}

func (k *Kyber) SecretKey() (sK []byte, esK string, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("[Kyber] Error generating secret key")
		}
	}()

	cT := make([]byte, pke.CiphertextSize)
	pT := make([]byte, pke.PlaintextSize)
	seed := make([]byte, pke.EncryptionSeedSize)

	if _, err := io.ReadFull(rand.Reader, pT); err != nil {
		return nil, "", err
	}

	if _, err := io.ReadFull(rand.Reader, seed); err != nil {
		return nil, "", err
	}

	hasher := sha3.New256()
	hasher.Write(pT)
	hpT := hasher.Sum(nil)

	k.publicKey.EncryptTo(cT, hpT, seed)

	return hpT, base64.StdEncoding.EncodeToString(cT), nil
}
