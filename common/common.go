package common

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	// DateFormat 時刻のフォーマット
	DateFormat = "20060102_1504"
	// DirFormat ディレクトリ名のフォーマット
	DirFormat = DateFormat + "_%s"
	// FileFormat  ファイル名のフォーマット
	FileFormat = DateFormat + "_%s.txt"
)

// CreateEmptyFile 空ファイルを作成する
func CreateEmptyFile(name string) error {
	prefix := time.Now().Format(FileFormat)
	f, err := os.OpenFile(fmt.Sprintf(prefix, name), os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

// CreateDir ディレクトリを作成する
func CreateDir(name string) error {
	prefix := time.Now().Format(DirFormat)
	return os.Mkdir(fmt.Sprintf(prefix, name), 0755)
}

// UpdateTimePath パス名を更新する
func UpdateTimePath(oldPath string) error {
	if _, err := os.Stat(oldPath); err != nil {
		return err
	}

	parent, base := filepath.Split(oldPath)

	if !isFormatedPath(base) {
		return fmt.Errorf("can't update path [%s]", oldPath)
	}

	newPath := parent + updateBasename(base)

	if oldPath == newPath {
		return nil
	}
	return os.Rename(oldPath, newPath)
}

func isFormatedPath(basename string) bool {
	if len(basename) < len(DateFormat) {
		return false
	}
	if _, err := time.Parse(DateFormat, basename[:len(DateFormat)]); err != nil {
		return false
	}
	return true
}
func updateBasename(basename string) string {
	prefix := time.Now().Format(DateFormat)
	return prefix + basename[len(DateFormat):]
}
