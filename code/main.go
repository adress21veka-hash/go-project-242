package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {

	info, err := os.Stat(path)

	if err != nil {
		return "", errors.New("Ошибка в определении это файл или папка")
	}

	if !info.IsDir() {
		return sizeOfFile(info), nil
	}

	if info.IsDir() {
		return sizeOfSumFiles(path), nil
	}

	return "", errors.New("Ошибка, ничего не подошло под условия")
}

func sizeOfFile(info os.FileInfo) string {
	return fmt.Sprintf("%dB", info.Size())
}

func sizeOfSumFiles(path string) string {

	filesindirectory := getFilesInDirectory(path)
	return sizeOfFiles(filesindirectory)

}

func getFilesInDirectory(path string) []string {
	var files []string

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, filepath.Join(path, e.Name()))
		}
	}

	return files
}

func sizeOfFiles(files []string) string {
	var total int64
	for _, f := range files {
		info, err := os.Stat(f)
		if err != nil {
			continue
		}
		total += info.Size()
	}
	return fmt.Sprintf("%dB", total)
}
