package graph

import (
	"fmt"
	"log"
)

// isSafe проверяет, можно ли добавить вершину в гамильтонов цикл
func (g *Graph) isSafe(v int, pos int, path []int) bool {
	// Проверяем, есть ли ребро между добавляемой вершиной и последней вершиной в пути
	if g.Amatrix[path[pos-1]][v] == 0 {
		return false
	}

	// Проверяем, не была ли добавляемая вершина уже включена в путь
	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}

	return true
}

// hamiltonCycle использует метод "проб и ошибок" для поиска гамильтонова цикла
func (g *Graph) hamiltonCycle(path []int, pos int) bool {
	// Базовый случай: если все вершины включены в цикл
	if pos == len(g.Amatrix) {
		// И последняя вершина связана с первой вершиной
		if g.Amatrix[path[pos-1]][path[0]] == 1 {
			return true
		} else {
			return false
		}
	}

	// Пытаемся добавить каждую вершину в цикл
	for v := 1; v < len(g.Amatrix); v++ {
		// Проверяем, можно ли добавить эту вершину
		if g.isSafe(v, pos, path) {
			path[pos] = v

			// Рекурсивно вызываем hamiltonCycle для следующей позиции
			if g.hamiltonCycle(path, pos+1) {
				return true
			}

			// Если добавление вершины v не приводит к решению,
			// то удаляем ее из цикла
			path[pos] = -1
		}
	}

	return false
}

// hamiltonian инициирует поиск гамильтонова цикла
func (g *Graph) Hamiltonian() []int {

	Am := copy2DSlice(g.Amatrix)

	path := make([]int, len(g.Amatrix))

	// Инициализируем путь как -1
	for i := range path {
		path[i] = -1
	}

	// Начинаем с вершины 0
	path[0] = 0
	if !g.hamiltonCycle(path, 1) {
		g.Amatrix = Am
		return []int{}
	}

	//добавляем начальную вершину в путь чтобы замкнуть
	path = append(path, path[0])
	g.Amatrix = Am
	return path
}

/*
алгоритм начинает с вершины 0 и пробует включить каждую вершину в гамильтонов цикл.
Если вершина может быть добавлена, то она добавляется, и рекурсивно вызывается hamiltonCycle
для следующей вершины. Если вершина не может быть добавлена в гамильтонов цикл,
то она удаляется, и алгоритм переходит к следующей вершине.
Если все вершины пробовались и ни одна не может быть добавлена, то возвращается false.
Если все вершины были включены в цикл и последняя вершина связана с первой вершиной,
то возвращается true, и мы нашли гамильтонов цикл.
*/

func (G *Graph) GamiltonTransform() {

	log.Print("func (G *Graph) EilerTransform() ")
	log.Print("делаем граф неориентированным")
	G.OrientToUnoriet()

	/*
		пока index, минимальный элемент из списка степеней вершин < n/2 {

			добавить столько вершин, чтобы его степень была n/2
		}
	*/
	// список вершин с нечетной степенью
	oddV := listOddV(G)

	fmt.Printf("список вершин с нечетными степенями %d \n", oddV)
	G.PrintLabel("граф:")

	pairs := getPairs(oddV, 0)

	for _, pair := range pairs {
		if G.Amatrix[pair[0]][pair[1]] == 0 {

			FlagP0inOddV := false
			FlagP1inOddV := false

			for _, v := range oddV {
				if v == pair[0] {
					FlagP0inOddV = true
				}
				if v == pair[1] {
					FlagP1inOddV = true
				}
			}

			//если оба значения есть в списке не четных
			if FlagP0inOddV && FlagP1inOddV {
				G.Amatrix[pair[0]][pair[1]] = 1
				G.Amatrix[pair[1]][pair[0]] = 1

				G.BandwidthMatrix[pair[0]][pair[1]] = 1
				G.BandwidthMatrix[pair[1]][pair[0]] = 1

				//удаляем вершины из списка не четных
				for i, v := range oddV {
					if v == pair[0] || v == pair[1] {
						oddV[i] = -1
					}

				}

				fmt.Printf("добавленая связь %d - %d\n", pair[0], pair[1])
			}

		} else {
			continue
		}
	}

}
