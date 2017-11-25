package main

import (
	"log"
	"dataset"
)

func main() {

	x := dataset.DataNode{
		Name: "data 1",
		Properties: []dataset.DataProperty{
			{Name:"property 1", Value: 3.11},
		},
	}

	log.Println(x)

}
