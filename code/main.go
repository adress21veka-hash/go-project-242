package code

import (
	"errors"
	"fmt"
	"os"
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
		return fmt.Sprintf("%d B", size), nil
	}

	if info.IsDir() {
		size, err := sizeOfSumFiles(path, all)
		if err != nil {
			return "", err
		}
		if human {
			return convertToHuman(size), nil
		}
		return fmt.Sprintf("%d B", size), nil
	}

	return "", errors.New("ошибка, ничего не подошло под условия")
}

func sizeOfSumFiles(path string, all bool) (int64, error) {
	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		// если all=false — пропускаем скрытые (начинаются с точки)
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		total += info.Size()
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
		{"B", 1},
	}

	for _, u := range units {
		if bytes >= u.size && bytes%u.size == 0 {
			return fmt.Sprintf("%d %s", bytes/u.size, u.name)
		}
	}

	return fmt.Sprintf("%d B", bytes)
}
