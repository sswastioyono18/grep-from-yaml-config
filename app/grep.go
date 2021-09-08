package app

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	zapLogger "github.com/sswastioyono18/grep-from-yaml-config/log"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
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
		zapLogger.Logger.Info(fmt.Sprintf("Getting list of key from this file: %s \n", filePath))
		yamlFile, err := ioutil.ReadFile(path.Join(y.Project.RootPath, app, filePath))

		if err != nil {
			zapLogger.Logger.Fatal("Error", zap.Error(err))
		}

		if !strings.Contains("secret", filePath) {
			// to remove line break in yaml config
			zapLogger.Logger.Info(fmt.Sprintf("Removing line break from this file: %s", filePath))
			yamlFile = bytes.Replace(yamlFile, []byte("|-"), []byte(""), -1)
		}

		yamlData := make(map[string]interface{})
		err2 := yaml.Unmarshal(yamlFile, &yamlData)

		if err2 != nil {
			zapLogger.Logger.Fatal("Error", zap.Error(err2))
		}

		for _, v := range y.TracePath {
			yamlData = yamlData[v].(map[string]interface{})
		}

		used := true
		for key, _ := range yamlData {
			used = y.Grep(key, app)
		}
		if used {
			zapLogger.Logger.Info(fmt.Sprintf("😊 All keys are used in this file: %s", filePath))
		}
	}
}

func (y *YamlContent) Grep(key, app string) (used bool){
	used = false
	files, err := WalkMatch(filepath.Join(viper.GetString("ROOT_PATH"),app), "*.go")
	if err != nil {
		return
	}

	for _, fileName := range files {
		grep := exec.Command("grep", key, fileName)
		// Run and get the output of grep.
		res, _ := grep.Output()
		if string(res) != "" {
			used = true
			return
		}
	}

	if used == false {
		zapLogger.Logger.Info(fmt.Sprintf("😞 This key is not used: %s", key))
		return
	}

	return
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
			zapLogger.Logger.Info(fmt.Sprintf("👀 Scanning folder: %s", app))
			grepper.GetContent(app)
		}
}
