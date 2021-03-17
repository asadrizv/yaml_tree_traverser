package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"strings"
)

type NodeLineInfo struct {
	lineNum  int
	NodeData interface{}
}

func main() {

	lineDataMap := ParseLineData("deployment.yaml")
	fmt.Println("Node name with line data Map: ", lineDataMap)
	errLine := fetchLineNum("input.spec.spec.containers.name", lineDataMap)
	fmt.Println("error could be at lines/line: ", errLine)
}

func filterPossibleLines ( possibleLines []int, referredNodes []string, LineDataMap map[string][]int) []int{
	for i := len(referredNodes)-2; i >= 0; i-- {
		predecessorLines := LineDataMap[referredNodes[i]]
		for _, line := range predecessorLines{
			for _, possibleLine := range possibleLines{
				if possibleLine < line{
					possibleLines = remove(possibleLines, possibleLine)
				}
				if len(possibleLines) == 1{
					return possibleLines
				}
			}
		}
	}
	return possibleLines
}

func remove(s []int, r int) []int {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func fetchLineNum(errString string, LineDataMap map[string][]int) []int {
	referredNodes := strings.Split(errString, ".")
	lastNodeData := LineDataMap[referredNodes[len(referredNodes)-1]]
	if len(lastNodeData) > 1{
		lastNodeData = filterPossibleLines(lastNodeData, referredNodes, LineDataMap)
	}
	return lastNodeData
}

func ParseLineData(fileName string) map[string][]int {
	LineNumMap := make(map[string][]int)
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
			var processedNode string
			err = currentNode.Decode(&processedNode)
			if err != nil {
				fmt.Println("error in decoding node: ", err)
			}

			LineNumMap[processedNode] = append(LineNumMap[processedNode], currentNode.Line)
		}

		for _, child := range currentNode.Content {
			nodeQue = append(nodeQue, *child)
		}

	}
	return LineNumMap
}
