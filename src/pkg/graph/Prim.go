package graph

import (
	"math"
)

// Infinity constant
const Infinity = math.MaxInt32

// PrimAlgorithm - Метод для нахождения остовного дерева минимального веса с использованием алгоритма Прима.
func (g *Graph) PrimAlgorithm() [][]int {
	// Количество вершин в графе.
	n := len(g.Amatrix)
	// Инициализируем матрицу для остовного дерева.
	spanningTree := make([][]int, n)
	for i := range spanningTree {
		spanningTree[i] = make([]int, n)
	}

	// Инициализируем массив для отслеживания посещенных вершин.
	visited := make([]bool, n)
	// Посещаем начальную вершину.
	visited[0] = true

	// Запускаем цикл, который будет работать для n-1 ребра.
	for i := 1; i < n; i++ {
		min := math.MaxInt64
		x, y := 0, 0
		for j := 0; j < n; j++ {
			if visited[j] {
				for k := 0; k < n; k++ {
					if !visited[k] && g.Amatrix[j][k] != 0 {
						// Если k вершина еще не была посещена и вес ребра меньше минимума.
						if min > g.Amatrix[j][k] {
							// Обновляем минимум.
							min = g.Amatrix[j][k]
							x = j
							y = k
						}
					}
				}
			}
		}
		// Добавляем найденное минимальное ребро в остовное дерево.
		spanningTree[x][y] = g.Amatrix[x][y]
		spanningTree[y][x] = g.Amatrix[y][x]
		// Помечаем вершину как посещенную.
		visited[y] = true
	}
	// Возвращаем матрицу остовного дерева.
	return spanningTree
}

//// PrimAlgorithm function implements Prim's algorithm
//func (G *Graph) PrimAlgorithm() [][]int {
//	n := G.vCount
//	parent := make([]int, n)  // Array to store constructed MST
//	key := make([]int, n)     // Key values used to pick minimum weight edge in cut
//	mstSet := make([]bool, n) // To represent set of vertices not yet included in MST
//
//	// Initialize all keys as INFINITE
//	for i := 0; i < n; i++ {
//		key[i] = Infinity
//	}
//
//	// Always include first 1st vertex in MST.
//	key[0] = 0
//	parent[0] = -1 // First node is always root of MST
//
//	for count := 0; count < n-1; count++ {
//		// Pick the minimum key vertex from the
//		// set of vertices not yet included in MST
//		u := minKey(key, mstSet)
//
//		// Add the picked vertex to the MST Set
//		mstSet[u] = true
//
//		// Update key value and parent index of
//		// the adjacent vertices of the picked vertex.
//		// Consider only those vertices which are not
//		// yet included in MST
//		for v := 0; v < n; v++ {
//			// graph[u][v] is non zero only for adjacent vertices of m
//			// mstSet[v] is false for vertices not yet included in MST
//			// Update the key only if graph[u][v] is smaller than key[v]
//			if G.Amatrix[u][v] != 0 && mstSet[v] == false && G.Amatrix[u][v] < key[v] {
//				parent[v] = u
//				key[v] = G.Amatrix[u][v]
//			}
//		}
//	}
//
//	// Generate MST by mapping the parent array to adjacency matrix
//	mst := make([][]int, n)
//	for i := range mst {
//		mst[i] = make([]int, n)
//	}
//	for i := 1; i < n; i++ {
//		mst[i][parent[i]] = G.Amatrix[i][parent[i]]
//		mst[parent[i]][i] = G.Amatrix[i][parent[i]]
//	}
//
//	return mst
//}

// A utility function to find the vertex with minimum key value,
// from the set of vertices not
// yet included in MST
func minKey(key []int, mstSet []bool) int {
	min := Infinity
	minIndex := -1

	for v := 0; v < len(key); v++ {
		if mstSet[v] == false && key[v] < min {
			min = key[v]
			minIndex = v
		}
	}

	return minIndex
}
