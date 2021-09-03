package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	dir, _ := os.Getwd()

	yfile, err := ioutil.ReadFile(dir+"/.infra/helm/dev/config.yaml")

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[string]interface{})

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	id  := data["config"].(map[string]interface{})
	content  := id["data"].(map[string]interface{})

	for k, _ := range content {
		//fmt.Printf("%s -> %s\n", k, v)
		grep := exec.Command("grep", "-r", k, dir+"/internal")

		// Run and get the output of grep.
		res, _ := grep.Output()
		if res == nil || string(res) == "" {
			fmt.Println("This key is not used:", k)
			continue
		}
	}
}

