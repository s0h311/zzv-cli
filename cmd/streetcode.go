package cmd

import (
	"zzv/cli/internal/config"
	"zzv/cli/internal/utils"

	"github.com/spf13/cobra"
)

func GetStreetcodeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "streetcode",
		Short: "Import streetcodes",
		Run: func(cmd *cobra.Command, args []string) {
			additionalCmd := config.GetStreetcodeCmd()

			utils.ExecuteCmdInDocker(*additionalCmd.PathInDocker, config.GetDockerContainerName(), additionalCmd.Cmds[0])
		},
	}

	return cmd
}
