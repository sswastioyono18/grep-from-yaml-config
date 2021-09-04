package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os/exec"
)

type YamlContent struct {
	directory string
}

func (y *YamlContent) GetContent() (yamlData map[string]interface{}) {
	yamlFile, err := ioutil.ReadFile(y.directory+"/.infra/helm/dev/config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	yamlData = make(map[string]interface{})
	err2 := yaml.Unmarshal(yamlFile, &yamlData)

	if err2 != nil {
		log.Fatal(err2)
	}

	yamlData = yamlData["config"].(map[string]interface{})
	yamlData  = yamlData["yamlData"].(map[string]interface{})

	return yamlData
}

func (y *YamlContent) Grep(key string) {
	grep := exec.Command("grep", "-r", key, y.directory+"/internal")

	// Run and get the output of grep.
	res, _ := grep.Output()
	if res == nil || string(res) == "" {
		fmt.Println("This key is not used:", key)
	}

}

func GrepFromFile (config IGrep) {
	content := config.GetContent()
	for k, _ := range content {
		//fmt.Printf("%s -> %s\n", k, v)
		config.Grep(k)

	}
}
