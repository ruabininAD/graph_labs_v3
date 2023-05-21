package graph

import (
	"fmt"
	"log"
)

func (G *Graph) Floid() ([][]int, [][]int, map[string][]int) {

	n := G.vCount

	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			log.Printf("%d : %d ", i, j)
			// Проверяем, что элемент не находится на диагонали и равен 0
			if i != j && G.Amatrix[i][j] == 0 {
				matrix[i][j] = inf
			} else {
				matrix[i][j] = G.Amatrix[i][j]
			}
		}
	}
	log.Printf("n, matrix  %d   --   %v", n, matrix)
	dist := make([][]int, n)        // матрица расстояний
	next := make([][]int, n)        // матрица следующих вершин на пути
	paths := make(map[string][]int) // словарь для хранения всех путей

	// инициализация матриц расстояний и следующих вершин
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = matrix[i][j]
			if matrix[i][j] != 0 {
				next[i][j] = j
			} else {
				next[i][j] = -1
			}
		}
	}

	// алгоритм Флойда-Уоршелла
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	// сохраняем все пути
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && next[i][j] != -1 {
				start := i
				end := j
				path := []int{start}
				for start != end {
					start = next[start][end]
					path = append(path, start)
				}
				paths[fmt.Sprintf("%d-%d", i, j)] = path
			}
		}
	}

	return dist, next, paths
}
