package cmd

import (
	"zzv/cli/internal/config"
	"zzv/cli/internal/utils"

	"github.com/spf13/cobra"
)

func GetMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			additionalCmd := config.GetMigrationCmd()

			utils.ExecuteCmdInDocker(*additionalCmd.PathInDocker, config.GetDockerContainerName(), additionalCmd.Cmds[0])
		},
	}

	return cmd
}
