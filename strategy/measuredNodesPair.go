package strategy

import "dataset"

type measuredNodesPair struct {
	node1       dataset.DataNode
	node2       dataset.DataNode
	metricValue float64
}