package cli

import (
	"context"
	fileProcess "github.com/barisaydogdu/FileBackupToolCLI/pkg/file"
	"github.com/spf13/cobra"
	"log"
)

var (
	sourceFile string
	targetFile string
	period     int64
)

type CLI struct {
	ctx         context.Context
	ctxCancel   context.CancelFunc
	cmd         *cobra.Command
	subCommands map[string]*cobra.Command
	file        *fileProcess.File
}

func NewCLI(ctx context.Context, cancel context.CancelFunc) *CLI {
	cli := &CLI{
		ctx:       ctx,
		ctxCancel: cancel,
		cmd: &cobra.Command{
			Use:   "backup",
			Short: "A tool to copy files from source to target at a specified period",
			Long: `This application copies a file from a source path to a target path at regular intervals.
You can specify the source file, target file, and the frequency (in seconds) for the copy operation.`,
		},
		subCommands: make(map[string]*cobra.Command),
	}
	cli.cmd.SetContext(ctx)

	cli.file = fileProcess.NewFile(ctx, cli.cmd)

	cli.subCommands["backup"] = &cobra.Command{
		Use:   "file",
		Short: "Backup file with frequency",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cli.file.BackupFileWithPeriod(sourceFile, targetFile, period); err != nil {
				cmd.PrintErrf("there si something error with backup file %v\n", err)
			}
		},
	}
	cli.subCommands["backup"].SetContext(ctx)

	cli.cmd.PersistentFlags().StringVar(&sourceFile, "source", "", "")
	cli.cmd.PersistentFlags().StringVar(&targetFile, "target", "", "")
	cli.cmd.PersistentFlags().Int64Var(&period, "period", 0, "seconds")
	cli.cmd.AddCommand(cli.subCommands["backup"])

	return cli
}

func (c *CLI) Execute() {
	if err := c.cmd.ExecuteContext(c.ctx); err != nil {
		log.Fatal(err)
	}

	if sourceFile == "" || targetFile == "" {
		c.cmd.PrintErr("source and target files are required\n")
		c.ctxCancel()
		return
	}

	if period == 0 {
		c.cmd.PrintErr("please specify a period\n")
		c.ctxCancel()
		return
	}
}
