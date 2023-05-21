package generator

import (
	"graph_labs/src/pkg/graph"
	"log"
)

func PrufferGenerator(PrufferCode []int) (*graph.Graph, error) {
	SList := pruferDecode(PrufferCode)

	/*	//показать сипски связности
		for i, row := range SList {
			fmt.Printf("%d: ", i)

			for _, v := range row {
				fmt.Printf(" %d ", v)
			}
			fmt.Println()
		}
	*/

	VCount := len(SList)
	log.Printf("func PrufferGenerator(%v int)\n", PrufferCode)

	G, _ := graph.NewGraph(VCount)

	G.Flags["unoriented"] = true

	for i, row := range SList { //проход по всем спискам связносити
		for _, v := range row { //проход по значениям в списке связносити

			G.Amatrix[i][v] = 1 //заполнение матрицы связности
			G.Amatrix[v][i] = 1

		}

	}

	G.PrintLabel("PrufferGenerator") //fixme
	return G, nil
}

func pruferDecode(pruferCode []int) [][]int {
	// Количество вершин в дереве
	vertexCount := len(pruferCode) + 2
	// Инициализация списка смежности
	adjList := make([][]int, vertexCount)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	// Вершины и их степени
	vertexDegree := make([]int, vertexCount)
	for i := 0; i < vertexCount; i++ {
		vertexDegree[i] = 1
	}

	// Увеличиваем степень вершин в коде Прюфера
	for _, vertex := range pruferCode {
		vertexDegree[vertex]++
	}

	// Обрабатываем вершины с степенью 1
	for i := 0; i < len(pruferCode); {
		for j := 0; j < vertexCount; j++ {
			if vertexDegree[j] == 1 {
				// Добавляем ребро между вершинами j и pruferCode[i]
				adjList[j] = append(adjList[j], pruferCode[i])
				adjList[pruferCode[i]] = append(adjList[pruferCode[i]], j)

				// Уменьшаем степень вершин
				vertexDegree[j]--
				vertexDegree[pruferCode[i]]--

				i++
				break
			}
		}
	}

	// Добавляем последнее ребро
	for i, degree := range vertexDegree {
		if degree == 1 {
			adjList[vertexCount-1] = append(adjList[vertexCount-1], i)
			adjList[i] = append(adjList[i], vertexCount-1)
			break
		}
	}

	return adjList
}
