package dataset

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
)

type DataProviderCsv struct {
	Nodes []DataNode
}

const colNodeId = 1
const colPropId = 0
const colPropValue = 2

const colPropMapId = 0
const colPropMapName = 1


//userId, movieId, rating, timestamp
func NewDataProviderCsv(filename string) (dp DataProviderCsv) {

	dp = DataProviderCsv{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return dp
	}
	defer file.Close()

	mappedPropNames :=  parsePropertiesNames("movies.csv")

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

		r, ok := mappedNodes[record[colNodeId]]

		if ok{
			value, err = strconv.ParseFloat(record[colPropValue], 64)
			if err == nil {
				r.Properties = append(r.Properties, DataProperty{record[colPropId], mappedPropNames[record[colPropId]],value*2})
			}
		} else {
			mappedNodes[record[colNodeId]] =
				&DataNode{
					Name: mappedPropNames[record[colNodeId]],
					Properties: []DataProperty{},
				}
		}

		lineCount += 1
	}

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
