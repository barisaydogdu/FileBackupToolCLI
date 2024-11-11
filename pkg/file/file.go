package file

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	ctx context.Context
	cmd *cobra.Command
}

func NewFile(ctx context.Context, cmd *cobra.Command) *File {
	return &File{ctx: ctx, cmd: cmd}
}

func (f *File) BackupFileWithPeriod(source, destination string, period int64) error {
	ticker := time.NewTicker(time.Second * time.Duration(period))

	go func() {
		for {
			select {
			case <-f.ctx.Done():
				ticker.Stop()
				f.cmd.Println("file backup stopped")
				return
			case <-ticker.C:
				if err := f.BackupFile(source, destination); err != nil {
					f.cmd.PrintErr(err)
					return
				}
			}
		}
	}()

	return nil
}

func (f *File) BackupFile(source, destination string) error {
	f.cmd.Println("backing up file", source, destination)
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
		mt, err := f.GetModeTime(destPath)
		if err != nil {
			return err
		}

		_, err2 := os.Stat(destPath)
		if os.IsNotExist(err2) || info.ModTime().After(*mt) {
			return f.CopyFile(path, destPath)
		}

		return nil
	})
}

func (f *File) CopyFile(sourceFile, targetFile string) error {
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

func (f *File) GetModeTime(path string) (*time.Time, error) {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("there is something error with get file info : %v", err)
	} else if err != nil && os.IsNotExist(err) {
		return &time.Time{}, nil
	}

	modificationTime := info.ModTime()

	return &modificationTime, nil

}
