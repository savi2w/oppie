package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/savi2w/oppie/config"
)

type RSA struct {
	publicKey *rsa.PublicKey
}

func New() (result *RSA, err error) {
	block, _ := pem.Decode([]byte(config.PublicKey))
	if block == nil {
		return nil, errors.New("failed to decode RSA public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch publicKey := publicKey.(type) {
	case *rsa.PublicKey:
		return &RSA{
			publicKey: publicKey,
		}, nil
	default:
		return nil, fmt.Errorf("unknown public key type")
	}
}

func (s *RSA) Byte(message []byte) (result []byte, err error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, s.publicKey, message, nil)
}

func (s *RSA) Base64(message []byte) (result string, err error) {
	r, err := s.Byte(message)
	if err != nil {
		return string(r), err
	}

	return base64.StdEncoding.EncodeToString(r), nil
}
