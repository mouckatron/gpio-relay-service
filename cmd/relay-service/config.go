package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var conf config
var confFile = "/etc/relay-service.yml"

type config struct {
	Host   string      `yaml:"host"`
	Port   string      `yaml:"port"`
	Relays []confRelay `yaml:"relays" json:"relays"`
}

func (c *config) write() {
	data, err := yaml.Marshal(c)

	if err != nil {
		log.Print(err)
		return
	}

	err = ioutil.WriteFile(confFile, data, 0644)
}

func getConfig() (conf config) {

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("Error reading %s", confFile)
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v", conf)

	return
}

func (c *config) setRelays(relays map[string]*relay) {

	c.Relays = make([]confRelay, 0)

	for _, r := range relays {
		c.Relays = append(c.Relays, confRelay{
			r.GPIO,
			r.Name,
		})
	}
}
