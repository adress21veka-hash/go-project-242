package main

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"code/code"
	"context"
	"log"
	"os"
	"strings"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "[path]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()

			if path == "" {
				return fmt.Errorf("нужен путь")
			}

			files, err := code.GetPathSize(path, true, true, true)

			if err != nil {
				return fmt.Errorf("файлы ошибку принесли")
			}

			fmt.Printf("%s	%s", files, path)
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Action: func(ctx context.Context, cmd *cli.Command, value bool) error {
					path := cmd.Args().First()

					if path == "" {
						return fmt.Errorf("нужен путь")
					}

					files, err := code.GetPathSize(path, true, true, true)

					if err != nil {
						return fmt.Errorf("файлы ошибку принесли")
					}

					fmt.Printf("%s	%s", convertToMb(files), path)
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func convertToMb(files string) string {
	files = strings.TrimSpace(files)

	runes := []rune(files)
	if len(runes) == 0 || string(runes[len(runes)-1]) != "B" {
		return "invalid"
	}

	bytes := int64(len(runes) - 1)

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
