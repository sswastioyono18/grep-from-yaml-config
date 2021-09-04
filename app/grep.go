package app

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type YamlContent struct {
	Project Project
	TracePath []string
}

type Project struct {
	AppTarget     []string
	GrepDirSource []string
	GrepExtTarget string
	RootPath      string
}

func (y *YamlContent) GetContent(app string) () {
	for _, filePath := range y.Project.GrepDirSource {
		fmt.Println("getting key from this file ", filePath)
		yamlFile, err := ioutil.ReadFile(path.Join(y.Project.RootPath, app, filePath))

		if err != nil {
			log.Fatal(err)
		}

		yamlData := make(map[string]interface{})
		err2 := yaml.Unmarshal(yamlFile, &yamlData)

		if err2 != nil {
			log.Fatal(err2)
		}

		for _, v := range y.TracePath {
			yamlData = yamlData[v].(map[string]interface{})
		}

		for key, _ := range yamlData {
			y.Grep(key, app)
		}
	}
}

func (y *YamlContent) Grep(key, app string) {
	used := false
	files, err := WalkMatch(filepath.Join(viper.GetString("ROOT_PATH"),app), "*.go")
	if err != nil {
		return
	}

	for _, fileName := range files {
		grep := exec.Command("grep", key, fileName)
		// Run and get the output of grep.
		res, _ := grep.Output()
		if string(res) != "" {
			//fmt.Println(fmt.Sprintf("%s %s %s", key, "used in : ", fileName))
			used = true
			return
		}
	}

	if used == false {
		fmt.Println("This key is not used:", key)
	}
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func StartGrepFromFile(grepper IGrep) {
		appList := viper.GetStringSlice("APP_TARGET")
		for _, app := range appList {
			log.Println(" scanning folder: ", app)
			grepper.GetContent(app)
		}
}
