package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type LineVal struct {
	lineNum  int
	NodeData interface{}
}

func main() {
	var LineNumMap []LineVal
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
	for len(nodeQue) > 0 {
		var currentNode yaml.Node
		currentNode, nodeQue = nodeQue[0], nodeQue[1:]

		if len(currentNode.Value) != 0 {
			lineNum := currentNode.Line
			fmt.Println("Node Value: ", currentNode.Value, ", Node line: ", currentNode.Line)
			var processedNode LineVal
			processedNode.lineNum = lineNum
			err = currentNode.Decode(&processedNode.NodeData)
			if err != nil{
				fmt.Println("error in decoding node: ", err)
			}
			LineNumMap = append(LineNumMap, processedNode)

		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	fmt.Println("Line Num Map", LineNumMap)
}
