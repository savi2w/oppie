package traversal

import (
	"github.com/savi2w/oppie/encrypt"
	"github.com/savi2w/oppie/storage"
)

type Traversal struct {
	Walker    *Walker
	Encryptor *Encryptor
}

func New(badger *storage.Badger, core *encrypt.Core) *Traversal {
	return &Traversal{
		Walker:    NewWalker(badger),
		Encryptor: NewEncryptor(badger, core),
	}
}
