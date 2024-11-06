package cli

import (
	"fmt"
	"os"

	fileProcess "github.com/barisaydogdu/FileBackupToolCLI/pkg/file"
	"github.com/spf13/cobra"
)

var (
	sourceFile string
	targetFile string
	period     int64

	rootCmd = &cobra.Command{
		Use:   "filecopy",
		Short: "A tool to copy files from source to target at a specified period",
		Long: `This application copies a file from a source path to a target path at regular intervals.
You can specify the source file, target file, and the frequency (in seconds) for the copy operation.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&sourceFile, "sourcefile", "", "")
	rootCmd.PersistentFlags().StringVar(&targetFile, "targetfile", "", "")
	rootCmd.PersistentFlags().Int64Var(&period, "period", 0, "")
	rootCmd.AddCommand(backupCommand)
}

var backupCommand = &cobra.Command{
	Use:   "backupfile",
	Short: "Backup file with frequency",
	Run: func(cmd *cobra.Command, args []string) {
		if err := fileProcess.BackupFilewithPeriod(sourceFile, targetFile, period); err != nil {
			fmt.Errorf("There si something error with backup file %v", err)
		}
	},
}
