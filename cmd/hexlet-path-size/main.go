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
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
