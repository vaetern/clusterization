package dataset

import (
	"path/filepath"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

type DataProviderIrisCsv struct {
	IrisData [][]float64
	IrisLabes []string
}

func NewDataProviderIrisCsv(filename string) (dp DataProviderIrisCsv)  {
	dp = DataProviderIrisCsv{}

	filePath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	irisData := make([][]float64, len(lines))
	irisLabels := make([]string, len(lines))
	for ii, line := range lines {
		vector := strings.Split(line, ",")
		label := vector[len(vector)-1]
		vector = vector[:len(vector)-1]
		floatVector := make([]float64, len(vector))
		for jj := range vector {
			floatVector[jj], err = strconv.ParseFloat(vector[jj], 64)
		}
		irisData[ii] = floatVector
		irisLabels[ii] = label
	}

	dp.IrisData = irisData
	dp.IrisLabes = irisLabels

	return dp
}

func (d DataProviderIrisCsv) AsNodes() []DataNode{

	var data []DataNode

	for i, vector := range d.IrisData{
		newNode := DataNode{strconv.Itoa(i), []DataProperty{}}
		for k, v := range vector  {
			newNode.Properties = append(newNode.Properties, DataProperty{strconv.Itoa(k),strconv.Itoa(k), v})
		}
		data = append(data, newNode)
	}

	return data
}