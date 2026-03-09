package services

import (
	"io"
	"os"
	"path/filepath"
)

func CopyDir(src string, dst string) error {

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, os.ModePerm)
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(target)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)

		return err
	})
}