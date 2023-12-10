package encrypt

import (
	"github.com/savi2w/zoldyck/encrypt/aes"
	"github.com/savi2w/zoldyck/encrypt/rsa"
)

type Core struct {
	rsa *rsa.RSA
}

func New() (e *Core, err error) {
	rsa, err := rsa.New()
	if err != nil {
		return nil, err
	}

	return &Core{rsa}, nil
}

func (c *Core) File(loc string) (fileKey string, err error) {
	return aes.File(loc, c.rsa)
}
