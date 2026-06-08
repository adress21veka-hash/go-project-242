package main

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"code/code"
	"context"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "[path]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return pathSize(ctx, cmd, false, false, false)
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Action: func(ctx context.Context, cmd *cli.Command, value bool) error {
					return pathSize(ctx, cmd, false, true, false)
				},
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
				Action: func(ctx context.Context, cmd *cli.Command, value bool) error {
					return pathSize(ctx, cmd, false, false, true)
				},
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
				Action: func(ctx context.Context, cmd *cli.Command, value bool) error {
					return pathSize(ctx, cmd, true, false, false)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func pathSize(ctx context.Context, cmd *cli.Command, recursive, human, all bool) error {
	path := cmd.Args().First()

	if path == "" {
		return fmt.Errorf("нужен путь")
	}

	files, err := code.GetPathSize(path, false, false, false)

	if err != nil {
		return fmt.Errorf("файлы ошибку принесли")
	}

	fmt.Printf("%s	%s", files, path)
	return nil
}
