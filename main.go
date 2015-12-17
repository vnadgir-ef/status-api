package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Config struct {
	HttpPort int `yaml:"httpPort"`
}

func readConfiguration() Config {
	var configFileName string
	flag.StringVar(&configFileName, "configFile", "config-local.yml", "Configuration file to use...")
	flag.Parse()

	filename, _ := filepath.Abs("./" + configFileName)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	log.Printf("Value: %v\n", config)

	return config
}

func main() {

	config := readConfiguration()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.HttpPort), router))
}
