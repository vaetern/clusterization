package dataset

type DataNode struct {
	Name       string
	Properties []DataProperty
}

type DataProperty struct {
	Name  string
	Value float64
}
