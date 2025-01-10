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

type Config struct {
	BaseDirEnv          string    `json:"baseDirEnv"`
	DockerContainerName string    `json:"dockerContainerName"`
	Projects            []Project `json:"projects"`
}

func getConfig() *Config {
	// GET PATH FROM ENV VARIABLE
	configJson, err := os.Open("~/zzv.config.json")

	if err != nil {
		fmt.Println(err)
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
