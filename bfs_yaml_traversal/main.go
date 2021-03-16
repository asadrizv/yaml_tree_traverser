package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	yamlFile, err := ioutil.ReadFile("deployment.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var nodeQue []yaml.Node
	var root yaml.Node

	err = yaml.Unmarshal(yamlFile, &root)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	nodeQue = append(nodeQue, *root.Content[0])
	var totalNodes = 0
	for len(nodeQue) > 0 {
		var currentNode yaml.Node
		currentNode, nodeQue = nodeQue[0], nodeQue[1:]

		if len(currentNode.Value) != 0 {
			fmt.Println("Node Value: ", currentNode.Value, ", Node line: ", currentNode.Line)
			totalNodes++
		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	fmt.Println(totalNodes)

}
