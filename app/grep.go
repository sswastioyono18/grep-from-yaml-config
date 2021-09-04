package app

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
)

type YamlContent struct {
	FileTarget string
	Project Project
	TracePath []string
}

type Project struct {
	AppTarget     []string
	GrepDirSource string
	GrepDirTarget string
	RootPath      string
}

func (y *YamlContent) GetContent(app string) (yamlData map[string]interface{}) {
	yamlFile, err := ioutil.ReadFile(path.Join(y.Project.RootPath, app, ".infra/helm/dev/config.yaml"))

	if err != nil {
		log.Fatal(err)
	}

	yamlData = make(map[string]interface{})
	err2 := yaml.Unmarshal(yamlFile, &yamlData)

	if err2 != nil {
		log.Fatal(err2)
	}

	for _, v := range y.TracePath {
		yamlData = yamlData[v].(map[string]interface{})
	}

	return yamlData
}

func (y *YamlContent) Grep(key, app string) {
	searchPath := path.Join(y.Project.RootPath, app, y.Project.GrepDirTarget)
	grep := exec.Command("grep", "-r", key, searchPath)

	// Run and get the output of grep.
	res, _ := grep.Output()
	if res == nil || string(res) == "" {
		fmt.Println("This key is not used:", key)
	}
}

func StartGrepFromFile(grepper IGrep) {
	appList := viper.GetStringSlice("APP_TARGET")
	for _, app := range appList {
		log.Println(" scanning folder: ", app)
		content := grepper.GetContent(app)
		for k, _ := range content {
			grepper.Grep(k, app)
		}
	}
}
