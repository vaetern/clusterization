package strategy

import (
	"dataset"
	"metrics"
	"fmt"
	"sync"
)

type FindClosestNodeStrategy struct {
	Metric metrics.Metric
}

type measuredNodesPair struct {
	node1       dataset.DataNode
	node2       dataset.DataNode
	metricValue float64
}

func NewFindClosestNodeStrategy(m metrics.Metric) (strat *FindClosestNodeStrategy) {
	strat = &FindClosestNodeStrategy{Metric: m}

	return strat
}


func (s FindClosestNodeStrategy) Process(Nodes []dataset.DataNode) {

	wg := sync.WaitGroup{}

	workersNum := 8

	outCh := make(chan measuredNodesPair)

	step := int64(len(Nodes) / workersNum)
	offsetX := int64(0)
	offsetY := step

	for i := 1; i <= workersNum; i++ {

		wg.Add(1)
		go s.reportClosest(Nodes, offsetX, offsetY, outCh, &wg)

		offsetX = int64(step) * int64(i) + 1
		offsetY = int64(step) * (int64(i) + 1)

	}

	go func() {
		wg.Wait()
		close(outCh)
	}()




	var result []measuredNodesPair


	for pair := range outCh {
		if pair.metricValue == 1 {
			result = append(result, pair)
		}
	}


	for _,x := range result {
		fmt.Println("=", x.node1.Name, x.node2.Name, x.metricValue)
	}

}

func (s FindClosestNodeStrategy) reportClosest(Nodes []dataset.DataNode, offsetStart int64, offsetFinish int64, outCh chan<- measuredNodesPair, wg *sync.WaitGroup) {
	fmt.Println("->", "worker up", "{", offsetStart,"-",offsetFinish,"}")

	curDistance := .0
	for index, x := range Nodes {
		if int64(index) >= offsetStart && int64(index) <= offsetFinish {
			for _, y := range Nodes {
				if x.Name != y.Name {
					curDistance = s.Metric.Distance(x.Properties, y.Properties)
					if curDistance == 1 {
						outCh <- measuredNodesPair{x, y, curDistance}
					}
				}
			}
		}
	}

	fmt.Println("->", "worker X", "{", offsetStart,"-",offsetFinish,"}")
	defer wg.Done()
}

