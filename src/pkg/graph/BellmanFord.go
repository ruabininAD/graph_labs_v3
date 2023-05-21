package graph

import (
	"fmt"
	"math"
)

func (G *Graph) BellmanFord(start, end int) (int, []int, error) {
	// Инициализация массивов расстояний и пути
	graph := G.Amatrix
	n := len(graph)
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[start] = 0

	// Проход алгоритма по всем вершинам
	for i := 0; i < n-1; i++ {
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if graph[u][v] != 0 && dist[u]+graph[u][v] < dist[v] {
					dist[v] = dist[u] + graph[u][v]
					prev[v] = u
				}
			}
		}
	}

	// Проверка наличия отрицательных циклов
	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && dist[u]+graph[u][v] < dist[v] {
				return -1, nil, fmt.Errorf("Отрицательный цикл") // Отрицательный цикл
			}
		}
	}

	// Восстановление пути из начальной вершины в конечную
	path := []int{}
	u := end
	for u != -1 {
		path = append([]int{u}, path...)
		u = prev[u]
	}

	return dist[end], path, nil
}
