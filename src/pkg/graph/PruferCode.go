package graph

import "fmt"

// Convert adjacency matrix to adjacency list
func AdjacencyMatrixToList(matrix [][]int) [][]int {
	n := len(matrix)
	list := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				list[i] = append(list[i], j)
			}
		}
	}
	return list
}

func (G *Graph) PruferCode() []int {
	SpannindTree := copy2DSlice(G.KruskalAlgorithm())
	for y := 0; y < G.vCount; y++ {
		for x := y; x < G.vCount; x++ {
			if SpannindTree[y][x] != 0 {
				SpannindTree[x][y] = SpannindTree[y][x]
			}
		}
	}

	tree := AdjacencyMatrixToList(SpannindTree)

	//print spanning
	for i, r := range tree {
		fmt.Printf("%d: ", i)
		for _, v := range r {
			fmt.Printf(" %d ", v)
		}
		fmt.Println()
	}

	n := len(tree)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		degree[i] = len(tree[i])
	}

	prufer := make([]int, n-2)
	leaf := -1

	for i := 0; i < n-2; i++ {
		// Find the next leaf
		for j := 0; j < n; j++ {
			if degree[j] == 1 {
				leaf = j
				break
			}
		}

		// Find the neighbor of the leaf
		neighbor := -1
		for _, v := range tree[leaf] {
			if degree[v] > 0 {
				neighbor = v
				break
			}
		}

		prufer[i] = neighbor
		degree[leaf]--
		degree[neighbor]--
	}

	return prufer
}

//// Generate Prufer code from adjacency list
//func (G *Graph) PruferCode( /*tree [][]int*/) []int {
//	tree := AdjacencyMatrixToList(G.KruskalAlgorithm())
//	for i, r := range tree {
//		fmt.Printf("%d: ", i)
//		for _, v := range r {
//			fmt.Printf(" %d ", v)
//		}
//		fmt.Println()
//	}
//
//	n := len(tree)
//	degree := make([]int, n+1)
//	for i := 0; i < n; i++ {
//		degree[i] = len(tree[i])
//	}
//
//	ptr := 0
//	for degree[ptr] != 1 && ptr <= n {
//		ptr++
//	}
//
//	leaf := ptr
//	prufer := make([]int, n-2)
//
//	for i := 0; i < n-2; i++ {
//		for _, v := range tree[leaf] {
//			degree[v]--
//			if degree[v] == 1 && v < ptr {
//				prufer[i] = v
//				leaf = v
//				break
//			}
//		}
//		for ptr <= n && degree[ptr] != 1 {
//			ptr++
//		}
//		if i < n-2 && degree[ptr] == 1 {
//			prufer[i] = ptr
//			leaf = ptr
//		}
//	}
//	return prufer
//}

//
//func main() {
//	// Adjacency matrix
//	adjacencyMatrix := [][]int{
//		{0, 1, 0, 0, 0},
//		{1, 0, 1, 0, 0},
//		{0, 1, 0, 1, 1},
//		{0, 0, 1, 0, 0},
//		{0, 0, 1, 0, 0},
//	}
//	// Convert adjacency matrix to adjacency list
//	adjacencyList := AdjacencyMatrixToList(adjacencyMatrix)
//	fmt.Println("Adjacency List: ", adjacencyList)
//
//	// Generate Prufer code
//	pruferCode := PruferCode(adjacencyList)
//	fmt.Println("Prufer Code: ", pruferCode)
//}
