package cmd

import (
	"fmt"
	"zzv/cli/internal/utils"

	"github.com/spf13/cobra"
)

var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all projects",
	Run: func(cmd *cobra.Command, args []string) {
		pullAllProjects()
	},
}

func pullAllProjects() {
	baseDir := utils.GetEnv("ZZV_PROJECT_DIR")

	projects := []string{
		"zzv-core",
		"zzv-api",
		"zzv-frontend",
	}

	for _, project := range projects {
		projectDir := fmt.Sprintf("%s/%s", baseDir, project)

		utils.PrintlnColorful(utils.Magenta, fmt.Sprintf("Switching to main branch: %s", project))
		utils.ExecuteCmd("git", "-C", projectDir, "switch", "main")

		utils.PrintlnColorful(utils.Magenta, fmt.Sprintf("Pulling: %s", project))
		utils.ExecuteCmd("git", "-C", projectDir, "pull")

		utils.PrintDivider()
	}
}
