package util

import (
	module "golaer/modile"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var ConfigFile string
var Config = new(module.Yaml)

func InitConfig() {
	conf := new(module.Yaml)
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	Config = conf
}

