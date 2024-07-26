package main

import (
	"io"
	"os"
	"path/filepath"
)

// Реализуйте программу для синхронизации содержимого двух директорий.
func syncDirectories(sourceDir, destDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return copyFile(path, destPath)
	})
}

func copyFile(sourceFile, destFile string) error {
	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	srcInfo, err := os.Stat(sourceFile)
	if err != nil {
		return err
	}

	return os.Chmod(destFile, srcInfo.Mode())
}
