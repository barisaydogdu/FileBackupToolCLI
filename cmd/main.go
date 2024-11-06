package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	cli "github.com/barisaydogdu/FileBackupToolCLI/handlers/cli"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs //sinyal alındığında kanal tetiklenir
		fmt.Println("\nBackup process stopping")
		done <- true
	}()

	fmt.Println("The backup process has started. Press CTRL + C to exit.")

	for {
		select {
		case <-done:
			fmt.Println("Program stopped")
			return
		default:
			cli.Execute()
		}
	}

}
