package cmd

import (
	"os"

	"github.com/bubblegutz/wfs/fs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().SortFlags = false
}

var (
	startCmd = &cobra.Command{
		Use:   "mount <path-to-mountpoint>",
		Short: "Mount and start the Webfilesystem",
		Long:  `Let's mount the web to your machine.`,
		Run:   startWFSCmdRun,
		Args:  cobra.ExactArgs(1),
	}
)

func startWFSCmdRun(cmd *cobra.Command, args []string) {
	fs.NewFS(os.Stdout, args[0])
}
