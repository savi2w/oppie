package fsutil

import (
	"os"

	"github.com/0x9ef/go-wiper/wipe"
)

func Wipe(path string) error {
	// Check if DoD is done correctly, I don't trust this library
	if err := wipe.Wipe(path, wipe.RuleUsDod5220_22_M); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}
