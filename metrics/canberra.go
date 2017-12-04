package metrics

import (
	"dataset"
	"math"
)

type CanberraMetric struct{}

func (CanberraMetric) Distance(prop1 []dataset.DataProperty, prop2 []dataset.DataProperty) float64 {
	commonRatingCount := 0

	totalSum := .0

	for _, x := range prop1 {
		for _, y := range prop2 {
			if x.Id == y.Id {
				commonRatingCount++
				absMinus := math.Abs(x.Value - y.Value)
				sumAbs := math.Abs(x.Value) +  math.Abs(y.Value)
				totalSum += absMinus/sumAbs
			}
		}
	}

	if commonRatingCount == 0 || totalSum == 0{return 1}

	return 1-1/totalSum
}
