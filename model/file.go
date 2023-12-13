package model

import (
	"encoding/json"

	"github.com/savi2w/oppie/storage"
)

type File struct {
	Key    string          `json:"-"`
	Badger *storage.Badger `json:"-"`

	Path        string  `json:"Path"`
	IsDeleted   bool    `json:"IsDeleted"`
	IsEncrypted bool    `json:"IsEncrypted"`
	Kyber       *string `json:"Kyber"`
}

func (file *File) Commit() error {
	json, err := json.Marshal(file)
	if err != nil {
		return err
	}

	if err := file.Badger.Set(file.Key, string(json)); err != nil {
		return err
	}

	return nil
}
