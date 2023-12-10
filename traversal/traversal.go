package traversal

import (
	"github.com/savi2w/zoldyck/encrypt"
	"github.com/savi2w/zoldyck/storage"
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
