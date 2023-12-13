package traversal

import (
	"encoding/json"

	"github.com/savi2w/oppie/encrypt/sha"
	"github.com/savi2w/oppie/model"
	"github.com/savi2w/oppie/storage"
)

func createFile(path string, badger *storage.Badger) error {
	key, err := sha.Hexadecimal([]byte(path))
	if err != nil {
		return err
	}

	file := &model.File{
		Key:    key,
		Badger: badger,

		Path:        path,
		IsDeleted:   false,
		IsEncrypted: false,
		Kyber:       nil,
	}

	return file.Commit()
}

func unmarshalFile(fileKey string, badger *storage.Badger, value []byte) (result *model.File, err error) {
	file := model.File{
		Key:    fileKey,
		Badger: badger,
	}

	if err := json.Unmarshal(value, &file); err != nil {
		return nil, err
	}

	return &file, nil
}
