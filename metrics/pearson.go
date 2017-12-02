package metrics

import (
	"math"
	"dataset"
)

type PearsonMetric struct{}

func (PearsonMetric) Distance(prop1 []dataset.DataProperty, prop2 []dataset.DataProperty) float64 {

	commonRatingCount := 0

	sum1, sum2 := .0, .0
	sum1Sq, sum2Sq := .0, .0
	totalSum := .0

	for _, x := range prop1 {
		for _, y := range prop2 {
			if x.Id == y.Id {

				commonRatingCount++

				sum1 += float64(x.Value)
				sum2 += float64(y.Value)
				sum1Sq += math.Pow(float64(x.Value), 2)
				sum2Sq += math.Pow(float64(y.Value), 2)
				totalSum += float64(x.Value * y.Value)
			}
		}
	}

	if commonRatingCount == 0 {
		return .0
	}

	n := float64(commonRatingCount)

	num := totalSum - (sum1 * sum2 / n)
	den := math.Sqrt((sum1Sq - math.Pow(sum1, 2)/n) * (sum2Sq - math.Pow(sum2, 2)/n))

	if den == 0 {
		return .0
	}

	r := num / den

	return r
}
