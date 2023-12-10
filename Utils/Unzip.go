package Utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func unzip(filename, pathname string) error {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		targetPath := filepath.Join(pathname, file.Name)
		if file.FileInfo().IsDir() {
			// Dosya bir klasörse
			os.MkdirAll(targetPath, os.ModePerm)
		} else {
			// Dosya bir dosyaysa
			if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
				return err
			}

			targetFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, zippedFile); err != nil {
				return err
			}
		}
	}

	return nil
}