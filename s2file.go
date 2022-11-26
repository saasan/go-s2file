package s2file

import (
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Rename(oldpath, newpath string) error {
	err := os.Link(oldpath, newpath)
	if err != nil {
		return err
	}

	return os.Remove(oldpath)
}
