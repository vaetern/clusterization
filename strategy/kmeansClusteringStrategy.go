package strategy

import (
	"dataset"
	"github.com/bugra/kmeans"
	"log"
	"fmt"
)


type KmeansClusteringStrategy struct {
	dataProvider *dataset.DataProviderIrisCsv
}

func NewKmeansClusteringStrategy(data *dataset.DataProviderIrisCsv) (s KmeansClusteringStrategy) {

	s = KmeansClusteringStrategy{data}

	return s
}

func (s KmeansClusteringStrategy) Process() {
	threshold := 10
	// Best Distance for Iris is Canberra Distance
	labels, err := kmeans.Kmeans(s.dataProvider.IrisData, 3, kmeans.CanberraDistance, threshold)
	if err != nil {
		log.Fatal(err)
	}

	for ii, jj := range labels {
		fmt.Println(ii,jj)
	}
}
