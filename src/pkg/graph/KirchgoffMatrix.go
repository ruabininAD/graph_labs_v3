package graph

import "fmt"

func determinant(matrix [][]int, n int) int {
	if n == 1 {
		return matrix[0][0]
	} else if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	} else {
		det := 0
		for col := 0; col < n; col++ {
			submatrix := getSubmatrix(matrix, 0, col, n)
			a := 0
			if col%2 == 0 {
				a = 1
			} else {
				a = -1
			}
			det += a * matrix[0][col] * determinant(submatrix, n-1)
		}

		return det
	}

}

func getSubmatrix(matrix [][]int, row int, col int, n int) [][]int {
	submatrix := make([][]int, n-1)
	for i := range submatrix {
		submatrix[i] = make([]int, n-1)
	}
	newRow := 0
	for i := 0; i < n; i++ {
		if i == row {
			continue
		}
		newCol := 0
		for j := 0; j < n; j++ {
			if j == col {
				continue
			}
			submatrix[newRow][newCol] = matrix[i][j]
			newCol++
		}
		newRow++
	}
	return submatrix
}

func (G *Graph) Kirchhoff_Main() {
	AjM := copy2DSlice(G.Amatrix)
	for y := 0; y < G.vCount; y++ {
		for x := y; x < G.vCount; x++ {
			if G.Amatrix[y][x] != 0 {
				AjM[x][y] = G.Amatrix[y][x]
			}
		}
	}

	matrix_Kirchhoff := AdjacencyToKirchhoff(AjM)
	n := len(matrix_Kirchhoff)
	fmt.Println("\n--- Нахождение числа остовных деревьев в графе (матричная теорема Кирхгофа) ---\n")
	fmt.Println("\nМатрица Кирхгофа для дезориентированного графа:\n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d\t", matrix_Kirchhoff[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	if n > 2 {
		m := getSubmatrix(matrix_Kirchhoff, 0, 0, n)
		det := determinant(m, n-1)
		fmt.Printf("Число остовных деревьев в графе: %d\n", det)
	} else {
		fmt.Println("\n!!! Вершин меньше, чем 2.")
	}
	fmt.Println("\n-------------------------------------------------------------------------------\n")
}

// AdjacencyToKirchhoff transforms an adjacency matrix to a Kirchhoff matrix
func AdjacencyToKirchhoff(adjacencyMatrix [][]int) [][]int {
	// Create a new matrix filled with zeros
	n := len(adjacencyMatrix)
	kirchhoffMatrix := make([][]int, n)
	for i := range kirchhoffMatrix {
		kirchhoffMatrix[i] = make([]int, n)
	}

	// Step 1: fill the diagonal with the degrees of the nodes
	for i := 0; i < n; i++ {
		degree := 0
		for j := 0; j < n; j++ {
			degree += adjacencyMatrix[i][j]
		}
		kirchhoffMatrix[i][i] = degree
	}

	// Step 2: for each edge between i and j, set the value in the Kirchhoff matrix to -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if adjacencyMatrix[i][j] > 0 {
				kirchhoffMatrix[i][j] = -1
			}
		}
	}

	return kirchhoffMatrix
}
