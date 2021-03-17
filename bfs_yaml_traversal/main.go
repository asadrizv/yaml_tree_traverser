package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type NodeLineInfo struct {
	lineNum  int
	NodeData interface{}
}

func main() {
	LineNumData := ParseLineDataToStruct("deployment.yaml")
	LineNumDataMap := ParseLineDataToMap("deployment.yaml")
	LineParsedMap := ParseLineDataString("deployment.yaml")

	fmt.Println("Standard parsing of lineNum data: ", LineNumData)
	fmt.Println("Dictionary based parsing of lineNum data: ", LineNumDataMap)
	fmt.Println("Consolidated line-wise parsing of lineNum data: ", LineParsedMap)

}

func ParseLineDataString(fileName string) map[int]string {
	LineNumMap := make(map[int]string)
	yamlFile, err := ioutil.ReadFile(fileName)
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
			var nodeData string
			err = currentNode.Decode(&nodeData)
			if err != nil {
				fmt.Println("error in decoding node: ", err)
			}
			val, ok := LineNumMap[lineNum]
			if ok{
				LineNumMap[lineNum] = val + ": " +nodeData
			} else{
				LineNumMap[lineNum]+=nodeData
			}

		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	return LineNumMap
}

func ParseLineDataToMap(fileName string) map[int][]NodeLineInfo {
	LineNumMap := make(map[int][]NodeLineInfo)
	yamlFile, err := ioutil.ReadFile(fileName)
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
			var processedNode NodeLineInfo
			processedNode.lineNum = lineNum
			err = currentNode.Decode(&processedNode.NodeData)
			if err != nil {
				fmt.Println("error in decoding node: ", err)
			}

			LineNumMap[lineNum] = append(LineNumMap[lineNum], processedNode)
		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	return LineNumMap
}

func ParseLineDataToStruct(fileName string) []NodeLineInfo {
	var LineNumMap []NodeLineInfo
	yamlFile, err := ioutil.ReadFile(fileName)
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
			var processedNode NodeLineInfo
			processedNode.lineNum = lineNum
			err = currentNode.Decode(&processedNode.NodeData)
			if err != nil {
				fmt.Println("error in decoding node: ", err)
			}
			LineNumMap = append(LineNumMap, processedNode)
		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	return LineNumMap
}
