package traversal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/savi2w/oppie/storage"
)

type Walker struct {
	Badger    *storage.Badger
	WaitGroup sync.WaitGroup
}

func NewWalker(badger *storage.Badger) *Walker {
	return &Walker{
		Badger:    badger,
		WaitGroup: sync.WaitGroup{},
	}
}

const AccessDenied = "Access is denied."

func (w *Walker) Walk(dir string, ignore []string) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("[Walker] Error walking: %s\n", rec)
		}

		w.WaitGroup.Done()
	}()

	entries, err := os.ReadDir(dir)
	if err != nil {
		message := err.Error()
		if strings.Contains(message, AccessDenied) {
			return
		}

		fmt.Printf("[Walker] Error reading directory `%s`: %s\n", dir, message)
		return
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			w.WaitGroup.Add(1)

			go w.Walk(path, ignore)
			continue
		}

		info, err := entry.Info()
		if err != nil {
			fmt.Printf("[Walker] Error getting file `%s` info: %s\n", path, err.Error())
			continue
		}

		if info.Mode()&os.ModeSymlink != 0 {
			continue
		}

		ignoreFile := false

		for _, ig := range ignore {
			if strings.Contains(path, ig) {
				ignoreFile = true
				break
			}
		}

		if ignoreFile {
			continue
		}

		if err := createFile(path, w.Badger); err != nil {
			fmt.Printf("[Walker] Error registering file `%s`: %s\n", path, err.Error())
			continue
		}

		fmt.Printf("[Walker] Registered file `%s`\n", path)
	}
}
