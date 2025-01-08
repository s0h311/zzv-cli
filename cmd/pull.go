package cmd

import (
	"fmt"
	"strings"
	"sync"
	"zzv/cli/internal/config"
	"zzv/cli/internal/utils"

	"github.com/spf13/cobra"
)

var pullAll bool

func GetPullCmd() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull all projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects := []string{
				"zzv-core",
				"zzv-api",
				"zzv-frontend",
			}

			if pullAll {
				projects = append(projects, "zzv-app-check24", "zzv-desktop-check24", "zzv-mobile-check24", "vendor/zzv-end2end")
			}

			pullProjects(projects)
		},
	}

	pullCmd.Flags().BoolVarP(&pullAll, "all", "a", false, "Pull all projects")

	return pullCmd
}

func pullProjects(projects []string) {
	baseDir := config.GetBaseDir()

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, project := range projects {
		projectDir := fmt.Sprintf("%s/%s", baseDir, project)

		wg.Add(1)

		go pullProjectProcess(project, projectDir, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

func pullProjectProcess(project string, projectDir string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- utils.SprintfColorful(utils.Magenta, "%s: Switching to main branch", strings.ToUpper(project))
	ch <- utils.ExecuteCmd("git", "-C", projectDir, "switch", "main")

	ch <- utils.SprintfColorful(utils.Magenta, "%s: Pulling", strings.ToUpper(project))
	ch <- utils.ExecuteCmd("git", "-C", projectDir, "pull")

	ch <- utils.GetDivider()
}
