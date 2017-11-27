package main

import (
	"log"
	"dataset"
	"metrics"
)

func main() {

	x := dataset.DataNode{
		Name: "data 1",
		Properties: []dataset.DataProperty{
			{Name:"property 1", Value: 3.11},
			{Name:"property 2", Value: 4.23},
			{Name:"property 3", Value: 5.40},
			{Name:"property 4", Value: 4.05},
			{Name:"property 5", Value: 9.28},
			{Name:"property 6", Value: 2.56},
			{Name:"property 7", Value: 1.16},
			{Name:"property 8", Value: 7.91},
		},
	}

	y := dataset.DataNode{
		Name: "data 2",
		Properties: []dataset.DataProperty{
			{Name:"property 1", Value: 4.11},
			{Name:"property 2", Value: 5.23},
			{Name:"property 3", Value: 6.40},
			{Name:"property 4", Value: 5.05},
			{Name:"property 5", Value: 8.28},
			{Name:"property 6", Value: 3.56},
			{Name:"property 7", Value: 2.16},
			{Name:"property 8", Value: 8.91},
		},
	}

	log.Println(metrics.EuclidianDistance(x.Properties, y.Properties))
	log.Println(metrics.PearsonScore(x.Properties, y.Properties))

}
