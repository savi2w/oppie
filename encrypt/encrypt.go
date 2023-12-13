package encrypt

import (
	"github.com/savi2w/oppie/encrypt/aes"
	"github.com/savi2w/oppie/encrypt/kyber"
)

type Core struct {
	k *kyber.Kyber
}

func New() (e *Core, err error) {
	k, err := kyber.New()
	if err != nil {
		return nil, err
	}

	return &Core{k}, nil
}

func (c *Core) File(loc string) (esK string, err error) {
	return aes.Counter(loc, c.k)
}
