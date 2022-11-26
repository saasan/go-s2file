package s2file

import (
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// ファイルを上書きしないRename
func Rename(oldpath, newpath string) error {
	// ハードリンクを作成
	err := os.Link(oldpath, newpath)
	if err != nil {
		return err
	}

	// 元ファイルを削除
	return os.Remove(oldpath)
}
