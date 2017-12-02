package main

import (
	ds "dataset"
	strat "strategy"
	"metrics"
	"runtime"
)

const Processes  = 8

func main() {

	runtime.GOMAXPROCS(Processes)


	dataProvider := ds.NewDataProviderCsv("data.csv")

	euclMetric := metrics.EuclideanMetric{}

	strategy := strat.NewFindClosestNodeStrategy(euclMetric)

	strategy.Process(dataProvider.Nodes)


}
