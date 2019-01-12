package tools

import (
	"os"
	"path/filepath"
)

func Args0() string {
	name, err := os.Executable()
	if err != nil {
		return filepath.Base(os.Args[0])
	}

	link, err := filepath.EvalSymlinks(name)
	if err != nil {
		return filepath.Base(name)
	}

	return filepath.Base(link)
}
