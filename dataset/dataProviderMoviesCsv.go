package dataset

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
)

type DataProviderMoviesCsv struct {
	Nodes []DataNode
}

const colNodeId = 1
const colPropId = 0
const colPropValue = 2

const colPropMapId = 0
const colPropMapName = 1


//userId, movieId, rating, timestamp
func NewDataProviderMoviesCsv(filename string, fileprop string) (dp DataProviderMoviesCsv) {

	dp = DataProviderMoviesCsv{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return dp
	}
	defer file.Close()

	mappedPropNames :=  parsePropertiesNames(fileprop)

	reader := csv.NewReader(file)

	mappedNodes := make(map[string]*DataNode)
	var value float64

	reader.Comma = ','
	lineCount := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		conv,_ := strconv.Atoi(record[colNodeId])
		if conv < int(300){

		r, ok := mappedNodes[record[colNodeId]]

		if ok {
			value, err = strconv.ParseFloat(record[colPropValue], 64)
			if err == nil {
				r.Properties = append(r.Properties, DataProperty{record[colPropId], mappedPropNames[record[colPropId]], value})
			}
		} else {
			value, err = strconv.ParseFloat(record[colPropValue], 64)
			mappedNodes[record[colNodeId]] =
				&DataNode{
					Name:       mappedPropNames[record[colNodeId]],
					Properties: []DataProperty{{record[colPropId], mappedPropNames[record[colPropId]], value}},
				}
		}

		lineCount += 1
	}
	}

	fmt.Println("->", filename, ":", lineCount)

	for _, v :=range mappedNodes{
		dp.Nodes = append(dp.Nodes, *v)
	}

	return dp

}

func parsePropertiesNames(filename string) (mapping map[string]string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return mapping
	}
	defer file.Close()

	reader := csv.NewReader(file)

	mapping = make(map[string]string)

	reader.Comma = ','
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_, ok := mapping[record[colPropMapId]]

		if !ok{
			mapping[record[colPropMapId]] = record[colPropMapName]
		}

	}

	return mapping
}
