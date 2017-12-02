package dataset

type DataNode struct {
	Name       string
	Properties []DataProperty
}

type DataProperty struct {
	Id    string
	Name  string
	Value float64
}
