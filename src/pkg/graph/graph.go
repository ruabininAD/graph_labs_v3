package graph

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Graph struct {
	vCount          int
	eCount          int
	Amatrix         [][]int
	Flags           map[string]bool
	BandwidthMatrix [][]int
}

func NewGraph(VCount int) (*Graph, error) {

	if VCount < 1 {
		return nil, fmt.Errorf("Попытка создать граф с нулем веришн с помощью функции NewGreph\n")
	}

	G := Graph{vCount: VCount, Amatrix: make([][]int, VCount), Flags: make(map[string]bool), BandwidthMatrix: make([][]int, VCount)}

	for i := 0; i < VCount; i++ {
		G.Amatrix[i] = make([]int, VCount)
		G.BandwidthMatrix[i] = make([]int, VCount)
	}

	G.Flags["oriented"] = false
	G.Flags["unoriented"] = false
	G.Flags["tree"] = false
	G.Flags[""] = false
	G.Flags["acyclic"] = false

	return &G, nil
}

func (G *Graph) Set(row, col, value int) {
	if G.Flags["oriented"] {

		G.Amatrix[row][col] = value
		G.eCount++

	} else {
		fmt.Errorf("граф не ориентированный")
		log.Printf("попытка внести ориентированное ребро в неориентированный граф")
	}
}

func (G *Graph) SetUnOrientedE(row, col, value int) {
	if G.Flags["unoriented"] {

		G.Amatrix[row][col] = value
		G.Amatrix[col][row] = value

	} else {
		fmt.Errorf("граф ориентированный, нет флага unoriented")
		log.Printf("попытка внести не ориентированное ребро в ориентированный граф")
	}

}

func (G *Graph) Get(row, col int) int {
	return G.Amatrix[row][col]
}

func (G *Graph) Print() {

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {

			fmt.Printf("\t%d", G.Amatrix[i][j])

		}

		fmt.Println() // перевод строки
	}

	fmt.Println() //переход на новую строку
}

func (G *Graph) PrintLabel(text string) {
	fmt.Println(text)
	G.Print()
}

func (G *Graph) Render() {
	file, err := os.OpenFile("src/graph.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""
	for i := 0; i < G.vCount; i++ {
		text = ""
		for j := 0; j < G.vCount; j++ {

			text += fmt.Sprintf("%d, ", G.Amatrix[i][j])

		}
		file.WriteString(text + "\n")
	}

}

func (G *Graph) GetECount() int {
	return G.eCount
}

func (G *Graph) GetVCount() int {
	return G.vCount
}

func (G *Graph) SetRandomWeight(flag string) {

	G.Flags["weight"] = true

	if flag == "-" {
		G.Flags["negativeWeight"] = true
	}

	//	negative weight
	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			if G.Get(i, j) != 0 {
				if flag == "+" {
					G.Set(i, j, rand.Intn(100))
				} else {
					G.Set(i, j, rand.Intn(100)-22)
				}

			}
		}
	}

}

func (G *Graph) RemoveWeights() {

	G.Flags["weight"] = false
	G.Flags["negativeWeight"] = false

	//	negative weight
	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			if G.Get(i, j) != 0 {
				G.Set(i, j, 1)
			}
		}
	}
}

func (G *Graph) GetAMatrix() [][]int {
	return G.Amatrix
}

func (G *Graph) SetRandomBandwidthMatrix(n int) {
	G.Flags["BandwidthMatrix"] = true

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {

			if G.Get(i, j) != 0 {
				G.BandwidthMatrix[i][j] = rand.Intn(n) + 2
			}
		}
	}

}

func (G *Graph) PrintBandwidthMatrix() {

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {

			fmt.Printf("\t%d", G.BandwidthMatrix[i][j])

		}

		fmt.Println() // перевод строки
	}

	fmt.Println() //переход на новую строку
}

func (G *Graph) PrintLabelBandwidthMatrix(text string) {
	fmt.Println(text)
	G.PrintBandwidthMatrix()
}

func (G *Graph) PrintFordFalkerson() {
	Bmatrix := G.FordFalkerson()

	if Bmatrix == nil {
		return
	}
	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {

			fmt.Printf("\t%d", Bmatrix[i][j])

		}

		fmt.Println() // перевод строки
	}
}

func (G *Graph) PrintPrimAlgorithm() {
	ostovTree := G.PrimAlgorithm()

	for _, row := range ostovTree {
		for _, v := range row {
			fmt.Printf("\t%d", v)
		}

		fmt.Println()
	}
}

func (G *Graph) PrintKraskalaAlgorithm() {
	ostovTree := G.KruskalAlgorithm()

	for _, row := range ostovTree {
		for _, v := range row {
			fmt.Printf("\t%d", v)
		}

		fmt.Println()
	}
}

func max(arr []int) int {
	max := arr[0]
	for _, element := range arr {
		if element > max {
			max = element
		}
	}
	return max
}

func min(arr []int) int {
	min := max(arr)
	for _, element := range arr {
		if element < min && element != 0 {
			min = element
		}
	}
	return min
}

func (G *Graph) UnorientToOriet() {
	G.Flags["oriented"] = false
	G.Flags["unoriented"] = true

	for y := 0; y < G.vCount; y++ {
		for x := 0; x < G.vCount; x++ {

			if y >= x {
				G.Amatrix[y][x] = 0
				G.BandwidthMatrix[y][x] = 0
			}

		}
	}
}

func (G *Graph) OrientToUnoriet() {
	G.Flags["oriented"] = true
	G.Flags["unoriented"] = false

	for y := 0; y < G.vCount; y++ {
		for x := y; x < G.vCount; x++ {

			if G.Amatrix[y][x] != 0 {
				G.Amatrix[x][y] = G.Amatrix[y][x]
				G.BandwidthMatrix[x][y] = G.BandwidthMatrix[y][x]
			}

		}
	}
}
