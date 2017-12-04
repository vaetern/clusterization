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

	dataProvider := ds.NewDataProviderMoviesCsv("storage/data_small.csv","storage/movies_small.csv")

	usedMetric := metrics.PearsonMetric{}

	strategy := strat.NewFindClosestNodeStrategy(usedMetric)

	strategy.Process(dataProvider.Nodes)

}
