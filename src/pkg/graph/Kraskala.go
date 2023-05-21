package graph

import "sort"

type Edge struct {
	weight int
	start  int
	end    int
}

func (g *Graph) KruskalAlgorithm() [][]int {
	edges := make([]Edge, 0)

	// Переводим матрицу смежности в список ребер
	for i := 0; i < g.vCount; i++ {
		for j := i + 1; j < g.vCount; j++ {
			if g.Amatrix[i][j] != 0 {
				edges = append(edges, Edge{g.Amatrix[i][j], i, j})
			}
		}
	}

	// Сортируем ребра по весу
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parent := make([]int, g.vCount)
	// Каждая вершина в своем подмножестве
	for i := range parent {
		parent[i] = i
	}

	union := func(i, j int) {
		pi, pj := find(parent, i), find(parent, j)
		parent[pi] = pj
	}

	// Матрица смежности для остовного дерева
	result := make([][]int, g.vCount)
	for i := range result {
		result[i] = make([]int, g.vCount)
	}

	// Проходим по всем ребрам в порядке увеличения веса
	for _, e := range edges {
		if find(parent, e.start) != find(parent, e.end) {
			union(e.start, e.end)
			result[e.start][e.end] = e.weight
			//result[e.end][e.start] = e.weight
		}
	}

	return result
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}
