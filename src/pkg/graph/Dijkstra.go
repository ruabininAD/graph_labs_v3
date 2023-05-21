package graph

import (
	"fmt"
	"log"
	"math"
)

const inf = math.MaxInt32

func (G *Graph) Dijkstra(start, finish int) (len int, path []int, err error) {

	log.Printf("func (G *Graph) Dijkstra(%d, %d int) (%v int ,%v []int, %s error))\n", start, finish, len, path, err)

	if G.Flags["negativeWeight"] == true {

		return 0, nil, fmt.Errorf("Алгоритм Дейкстры не работает в  взвешеным графом ")
	}

	dist := make([]int, G.vCount)     // слайс для хранения расстояний от начальной вершины до остальных вершин
	visited := make([]bool, G.vCount) // слайс для отслеживания посещенных вершин
	prev := make([]int, G.vCount)     // слайс для хранения предыдущих вершин на кратчайшем пути

	for i := 0; i < G.vCount; i++ {
		dist[i] = int(^uint(0) >> 1) // устанавливаем бесконечное расстояние для всех вершин, кроме начальной
		visited[i] = false
		prev[i] = -1
	}
	dist[start] = 0 // расстояние от начальной вершины до самой себя равно 0

	for i := 0; i < G.vCount; i++ {
		u := -1
		// выбираем вершину с наименьшим известным расстоянием и помечаем ее как посещенную
		for j := 0; j < G.vCount; j++ {
			if !visited[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}
		visited[u] = true

		// обновляем расстояния до соседних вершин
		for v := 0; v < G.vCount; v++ {
			if G.Get(u, v) != 0 && dist[u]+G.Get(u, v) < dist[v] {
				dist[v] = dist[u] + G.Get(u, v)
				prev[v] = u
			}
		}
	}

	// восстанавливаем кратчайший путь от конечной вершины до начальной
	path = []int{}
	u := finish
	for u != -1 {
		path = append([]int{u}, path...)
		u = prev[u]
	}

	count := -1
	for _, _ = range path {
		count++
	}

	return count, path, nil
}
