package main

import (
	// Standard Library Imports
	"flag"
	"io/ioutil"

	// External Imports
	"github.com/go-yaml/yaml"
	"github.com/sirupsen/logrus"

	// Internal Imports
	"github.com/matthewhartstonge/attribute"
)

func getConfig() *attribute.Config {
	f, err := ioutil.ReadFile(".attribute.yaml")
	if err != nil {
		logrus.Fatal("Unable to find .attribute.yaml file")
	}

	config := &attribute.Config{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		logrus.WithError(err).Fatal("Unable to read yaml file, please check your indentation and syntax")
	}

	return config
}

func WriteToDisk(filename string, contents []byte) {
	ioutil.WriteFile(filename, contents, 0644)
}

func main() {
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	config := getConfig()

	if license := attribute.GenerateLicense(config); license != nil {
		WriteToDisk("LICENSE", license)
	}

	if notice := attribute.GenerateNotice(config); notice != nil {
		WriteToDisk("NOTICE", notice)
	}

	attributions := attribute.GenerateAttributions(config)
	WriteToDisk("ATTRIBUTIONS.md", attributions)
}
