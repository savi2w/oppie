package osutil

import "os"

func GetSystemDrive() string {
	drive := os.Getenv("SystemDrive")
	if len(drive) != 0 {
		return drive + "\\"
	}

	return "C:" + "\\"
}

func GetIgnoreFolders() []string {
	vars := []string{
		"SystemRoot",
		"ProgramFiles",
		"ProgramFiles(x86)",
		"ProgramData",
		"AppData",
		"LocalAppData",
		"Public",
		"AllUsersProfile",
		"Temp",
		"CommonProgramFiles",
	}

	// Maybe this preset should be removed on walk time or something like that
	folders := []string{
		".lnk",
		".url",
		".dat",
		"\\Users\\Default\\",
		".log",
		".log2",
		".tmp",
		".sys",
		"\\NTUSER",
		".ini",
		"\\Searches\\",
		"\\$Recycle.Bin\\",
		"\\go\\",
		"\\.rustup\\",
		"\\.vscode\\",
		"\\Riot Games\\",
		"\\Repos\\",
		"\\Games\\",
		"\\.cargo\\",
		"\\.GamingRoot",
	}

	for _, key := range vars {
		value := os.Getenv(key)

		if len(value) != 0 {
			folders = append(folders, value)
		}
	}

	return folders
}
