package strategy

import (
	"metrics"
	"dataset"
	"sync"
	"fmt"
	"math"
	"image"
	"image/color"
	"os"
	"image/gif"
	"image/color/palette"
	"math/rand"
	"time"
)

type Draw2dClustersStrategy struct {
	Metric metrics.Metric
}

type point struct {
	x float64
	y float64
}

func NewDraw2dClustersStrategy(metric metrics.Metric) (s Draw2dClustersStrategy) {

	s = Draw2dClustersStrategy{metric}

	return s
}

func (s Draw2dClustersStrategy) Process(Nodes []dataset.DataNode) {

	rate := 0.0005

	var timelapse [][]point
	fmt.Println("->data population start")
	realdist := s.populateData(Nodes)
	fmt.Println("->data population done")

	n := len(Nodes)

	rand.Seed(time.Now().UTC().UnixNano())
	var loc []point
	for i := 0; i < n; i++ {
		loc = append(loc, point{rand.Float64(), rand.Float64()})
	}

	var fakedist [][]float64
	for i := 0; i < n; i++ {
		var row []float64
		for j := 0; j < n; j++ {
			row = append(row, .0)
		}
		fakedist = append(fakedist, row)
	}

	lasterror := 99999999999999999.1

	for m := 0; m < 10; m++ {
		fmt.Println("->iteration", m)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fakedist[i][j] = math.Sqrt(math.Pow(loc[i].x-loc[j].x, 2) + math.Pow(loc[i].y-loc[j].y, 2))
			}
		}

		var grad [][]float64
		for i := 0; i < n; i++ {
			var row []float64
			for j := 0; j < n; j++ {
				row = append(row, .0)
			}
			grad = append(grad, row)
		}

		totalerror := .0
		errorterm := .0

		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				if j == k {
					continue
				}
				if (*realdist)[j][k] != 0 {
					errorterm = (fakedist[j][k] - (*realdist)[j][k]) / (*realdist)[j][k]
				} else {
					errorterm = fakedist[j][k]
				}

				grad[k][0] += ((loc[k].x - loc[j].x) / fakedist[j][k]) * errorterm
				grad[k][1] += ((loc[k].y - loc[j].y) / fakedist[j][k]) * errorterm
				totalerror += math.Abs(errorterm)
			}
		}


		fmt.Println("approx:", totalerror )
		if lasterror < totalerror {
			break
		}
		lasterror = totalerror

		for k := 0; k < n; k++ {
			loc[k].x = - rate * grad[k][0]
			loc[k].y = - rate * grad[k][1]
		}
		timelapse = append(timelapse, loc)
	}

	createAnimatedImage(&timelapse)
}

func (s Draw2dClustersStrategy) populateData(Nodes []dataset.DataNode) *[][]float64 {
	wg := sync.WaitGroup{}

	workersNum := 8

	outCh := make(chan measuredNodesPair)

	step := int64(len(Nodes) / workersNum)
	offsetX := int64(0)
	offsetY := step

	for i := 1; i <= workersNum; i++ {

		wg.Add(1)
		go s.buildPairedDistance(Nodes, offsetX, offsetY, outCh, &wg)

		offsetX = int64(step)*int64(i) + 1
		offsetY = int64(step) * (int64(i) + 1)

	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	var result []measuredNodesPair

	for pair := range outCh {
		result = append(result, pair)
	}

	//for _, x := range result {
	//	fmt.Println(x.node1.Name, "|", x.node2.Name, x.metricValue)
	//}

	var distanceMatrix [][]float64
	{
	}

	for _, nodeA := range Nodes {
		var nodeArow []float64
		for _, nodeB := range Nodes {
			for _, pair := range result {
				if (nodeA.Name == pair.node1.Name && nodeB.Name == pair.node2.Name) ||
					(nodeA.Name == pair.node2.Name && nodeB.Name == pair.node1.Name) {
					nodeArow = append(nodeArow, pair.metricValue)
				}
			}
		}
		distanceMatrix = append(distanceMatrix, nodeArow)
	}

	return &distanceMatrix
}

func (s Draw2dClustersStrategy) buildPairedDistance(Nodes []dataset.DataNode, offsetStart int64, offsetFinish int64, outCh chan<- measuredNodesPair, wg *sync.WaitGroup) {
	fmt.Println("-w", offsetStart, "O")

	curDistance := .0
	for ix, x := range Nodes {
		if int64(ix) >= offsetStart-1 && int64(ix) <= offsetFinish+1 {
			for iy, y := range Nodes {
				if int64(iy) >= offsetStart-1 {
					curDistance = (math.Floor(1000*s.Metric.Distance(x.Properties, y.Properties)))/1000
					outCh <- measuredNodesPair{x, y, 10*(1 - curDistance)}
				}
			}
		}
	}

	fmt.Println("-w", offsetStart, "X")
	defer wg.Done()
}

func createAnimatedImage(timelapse *[][]point) {

	var w, h, midpoint int = 1001, 1001, 501
	steps := len(*timelapse)

	var value, topValue float64 = .0, .0

	for _, u := range (*timelapse)[len(*timelapse)-1]{
		value = math.Max(math.Abs(u.x),math.Abs(u.y))
		if value > topValue{topValue = value}
	}

	relCoef := float64(midpoint)/topValue


	fmt.Println("magnitude:", relCoef)
	fmt.Println("steps:", steps)

	var images []*image.Paletted
	var delays []int

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette.WebSafe)
		images = append(images, img)
		delays = append(delays, 4)
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				for _, z := range (*timelapse)[step] {
					if int(math.Floor(float64(midpoint) + z.x*relCoef)) == x && int(float64(midpoint) + math.Floor(z.y*relCoef)) == y {
						img.Set(x, y, color.White)
					}
				}
			}
		}
	}

	f, err := os.OpenFile("timelapse.gif", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		LoopCount:steps,
		Image: images,
		Delay: delays,
	})

}
