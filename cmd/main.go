package main

import (
	"context"
	"fmt"
	"github.com/barisaydogdu/FileBackupToolCLI/handlers/cli"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-sigs:
			cancel()
		}
	}()

	fmt.Println("The backup program has started. Press CTRL + C to exit.")

	app := cli.NewCLI(ctx, cancel)
	app.Execute()

	select {
	case <-ctx.Done():
		return
	}
}
