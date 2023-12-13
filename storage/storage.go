package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type Badger struct {
	DB *badger.DB
}

const DatabaseFile = "OPP-BadgerDB-%d"

func New() (*Badger, error) {
	file := fmt.Sprintf(DatabaseFile, time.Now().Unix())
	path := filepath.Join(os.TempDir(), file)

	cli, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &Badger{
		DB: cli,
	}, nil
}

func (s *Badger) Close() error {
	return s.DB.Close()
}

func (s *Badger) Set(key, value string) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}
