package file

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func BackupFilewithPeriod(source, destination string, period int64) error {
	ticker := time.NewTicker(time.Duration(period))
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := BackupFile(source, destination)
				if err != nil {
					fmt.Errorf("there is something wrong with backupfile %v", err)
					return
				}
			}
		}
	}()
	defer ticker.Stop()
	done <- true
	return nil
}

func BackupFile(source, destination string) error {
	return filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(source, path)
		destPath := filepath.Join(destination, relPath)

		if info.IsDir() {
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				if err := os.Mkdir(destPath, os.ModePerm); err != nil {
					return fmt.Errorf("error creating directory %q: %v", destPath, err)
				}
			}
			return nil
		}
		if os.Stat(destPath); os.IsNotExist(err) || info.ModTime().After(GetModeTime(destPath)) {
			return CopyFile(path, destPath)
		}
		return nil
	})
}

func CopyFile(sourceFile, targetFile string) error {
	input, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	return err
}

func GetModeTime(path string) time.Time {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Errorf("There is something error with get file info : %v", err)
		return time.Time{}
	}

	modificationTime := info.ModTime()

	return modificationTime

}

func SleepwithSecond() {

}
