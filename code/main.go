package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", errors.New("ошибка в определении это файл или папка")
	}

	if !info.IsDir() {
		size := info.Size()
		if human {
			return convertToHuman(size), nil
		}
		return fmt.Sprintf("%dB", size), nil
	}

	if info.IsDir() {
		size, err := sizeOfSumFiles(path, recursive, all)
		if err != nil {
			return "", err
		}
		if human {
			return convertToHuman(size), nil
		}
		return fmt.Sprintf("%dB", size), nil
	}

	return "", errors.New("ошибка, ничего не подошло под условия")
}

func sizeOfSumFiles(path string, recursive, all bool) (int64, error) {
	var total int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if entry.IsDir() && recursive {
			subSize, err := sizeOfSumFiles(filepath.Join(path, entry.Name()), recursive, all)
			if err != nil {
				continue
			}
			total += subSize
		} else if !entry.IsDir() {
			total += info.Size()
		}

	}

	return total, nil
}

func convertToHuman(bytes int64) string {
	units := []struct {
		name string
		size int64
	}{
		{"EB", 1 << 60},
		{"PB", 1 << 50},
		{"TB", 1 << 40},
		{"GB", 1 << 30},
		{"MB", 1 << 20},
		{"KB", 1 << 10},
	}

	for _, u := range units {
		if bytes >= u.size {
			return fmt.Sprintf("%.1f%s", float64(bytes)/float64(u.size), u.name)
		}
	}

	return fmt.Sprintf("%dB", bytes)
}
