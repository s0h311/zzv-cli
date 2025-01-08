package cmd

import (
	"fmt"
	"strings"
	"sync"
	"zzv/cli/internal/config"
	"zzv/cli/internal/utils"

	"github.com/spf13/cobra"
)

var buildAll bool

func GetBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build all projects",
		Run: func(cmd *cobra.Command, args []string) {
			projects := config.GetProjects()

			if buildAll {
				projects = config.GetAllProjects()
			}

			buildProjects(projects)
		},
	}

	cmd.Flags().BoolVarP(&buildAll, "all", "a", false, "Build all projects")

	return cmd
}

func buildProjects(projects []config.Project) {
	fmt.Println(projects)

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, project := range projects {
		wg.Add(1)

		go buildProjectProcess(project, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

func buildProjectProcess(project config.Project, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, buildCmd := range project.BuildCmds {
		ch <- utils.SprintfColorful(utils.Magenta, "%s: Building", strings.ToUpper(project.Name))

		if project.Type == "php" {
			utils.ExecuteCmdInDocker(project.Path, buildCmd)
			continue
		}

		if project.Type == "node" {
			utils.ExecuteNpmCmd(project.Path, buildCmd)
		}

		ch <- utils.GetDivider()
	}
}
