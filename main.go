package main

import (
	ds "dataset"
	"runtime"
	strat "strategy"
	"metrics"
)

const Processes  = 8

func main() {

	runtime.GOMAXPROCS(Processes)

	//dataProvider := ds.NewDataProviderMoviesCsv("storage/data.csv","storage/movies.csv")
	dataProvider := ds.NewDataProviderIrisCsv("storage/iris.csv")
	usedMetric := metrics.PearsonMetric{}
	strategy := strat.NewDraw2dClustersStrategy(usedMetric)
	strategy.Process(dataProvider.AsNodes())

	//dataProvider := ds.NewDataProviderIrisCsv("storage/iris.csv")
	//strategy := strat.NewKmeansClusteringStrategy(&dataProvider)
	//strategy.Process()

}
