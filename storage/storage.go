package storage

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgraph-io/badger/v4"
)

type Badger struct {
	DB *badger.DB
}

const DatabaseFile = "OPP-BadgerDB"

func New() (*Badger, error) {
	path := filepath.Join(os.TempDir(), DatabaseFile)
	unique := strconv.FormatInt(time.Now().Unix(), 10)

	cli, err := badger.Open(badger.DefaultOptions(path + unique))
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
