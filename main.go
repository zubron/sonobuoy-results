package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Item is the central format for plugin results. Various plugin
// types can be transformed into this simple format and set at a standard
// location in our results tarball for simplified processing by any consumer.
type Item struct {
	Name     string            `json:"name" yaml:"name"`
	Status   string            `json:"status" yaml:"status"`
	Metadata map[string]string `json:"meta,omitempty" yaml:"meta,omitempty"`
	Details  map[string]string `json:"details,omitempty" yaml:"details,omitempty"`
	Items    []Item            `json:"items,omitempty" yaml:"items,omitempty"`
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "Sonobuoy results file")
	flag.Parse()

	if fileName == "" {
		fmt.Println("Please provide sonobuoy results file by using -f option")
		return
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var results Item
	err = yaml.Unmarshal(yamlFile, &results)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Plugin: %v\n", results.Name)
	fmt.Printf("Overall Result: %v\n", results.Status)
	for _, result := range results.Items {
		if result.Metadata["type"] == "node" {
			fmt.Printf("Result from daemonset node %q: %v\n", result.Name, result.Status)
		} else if result.Metadata["type"] == "file" {
			fmt.Printf("Result from file %q: %v\n", result.Name, result.Status)
		}
	}
}
