package metrics

import "dataset"

type Metric interface{
	Distance(prop1 []dataset.DataProperty, prop2 []dataset.DataProperty) float64
}