package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/barisaydogdu/FileBackupToolCLI/handlers/cli"
)

var (
	sourceFilePath string
	targetFilePath string
	frequency      int64
	interval       time.Duration
)

func main() {
	// flag.StringVar(&sourceFilePath, "sourcefile", "", "")
	// flag.StringVar(&targetFilePath, "targetfile", "", "")
	// flag.Int64Var(&frequency, "frequency", 1, "")
	// flag.Parse()

	interval = time.Duration(frequency) * time.Second
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs //sinyal alındığında kanal tetiklenir
		fmt.Println("\nBackup process stopping")
		done <- true
	}()

	fmt.Println("The backup process has started. Press CTRL C to exit.")

	for {
		select {
		case <-done:
			fmt.Println("Program stopped")
			return
		default:
			// err := file.BackupFile(sourceFilePath, targetFilePath)
			// if err != nil {
			// 	fmt.Errorf("There is something went error when backup file %v", err)
			// }
			// time.Sleep(time.Duration(interval))
			cli.Execute()
		}
	}

}
