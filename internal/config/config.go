package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"zzv/cli/internal/utils"
)

type Project struct {
	Name          string   `json:"name"`
	Short         string   `json:"short"`
	Path          string   `json:"path"`
	PathInDocker  *string  `json:"pathInDocker"`
	Type          string   `json:"type"`
	Env           string   `json:"env"`
	BuildCmds     []string `json:"buildCmds"`
	IsMainProject bool     `json:"isMainProject"`
}

type AdditionalCmd struct {
	Name         string   `json:"name"`
	PathInDocker *string  `json:"pathInDocker"`
	Type         string   `json:"type"`
	Env          string   `json:"env"`
	Cmds         []string `json:"cmds"`
}

type Config struct {
	BaseDirEnv          string          `json:"baseDirEnv"`
	DockerContainerName string          `json:"dockerContainerName"`
	Projects            []Project       `json:"projects"`
	AdditionalCmds      []AdditionalCmd `json:"additionalCmds"`
}

func getConfig() *Config {
	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr != nil {
		fmt.Println(homeDirErr)
		os.Exit(1)
	}

	configJson, err := os.Open(fmt.Sprintf("%s/zzv.config.json", homeDir))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer configJson.Close()

	byteValue, _ := io.ReadAll(configJson)

	var config Config

	json.Unmarshal(byteValue, &config)

	return &config
}

func getBaseDirEnv() string {
	return getConfig().BaseDirEnv
}

func GetBaseDir() string {
	return utils.GetEnv(getBaseDirEnv())
}

func GetAllProjects() []Project {
	baseDirEnv := getBaseDirEnv()
	baseDir := GetBaseDir()
	projects := getConfig().Projects

	for index, project := range projects {
		projects[index].Path = strings.ReplaceAll(project.Path, "$"+baseDirEnv, baseDir)
	}

	return projects
}

func GetProjects() []Project {
	projects := []Project{}

	for _, project := range GetAllProjects() {
		if project.IsMainProject {
			projects = append(projects, project)
		}
	}

	return projects
}

func GetDockerContainerName() string {
	return getConfig().DockerContainerName
}

func GetMigrationCmd() AdditionalCmd {
	for _, cmd := range getConfig().AdditionalCmds {
		if cmd.Name == "migrate" {
			return cmd
		}
	}

	panic("Cannot find migrate command")
}
