package metrics

import (
	"math"
	"dataset"
)

type EuclideanMetric struct{}

func (EuclideanMetric) Distance(prop1 []dataset.DataProperty, prop2 []dataset.DataProperty) float64 {

	ratingDistanceRaw := .0

	for _, x := range prop1 {
		for _, y := range prop2 {
			if x.Id == y.Id {
				ratingDistanceRaw = ratingDistanceRaw + math.Pow(float64(x.Value-y.Value), 2)
			}
		}
	}

	euclidianDistance := 1 / (1 + math.Sqrt(ratingDistanceRaw))

	return euclidianDistance
}
