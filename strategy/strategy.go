package strategy

import (
	"dataset"
)

type DataProcessStrategy interface {
	Process(Nodes []dataset.DataNode)
}