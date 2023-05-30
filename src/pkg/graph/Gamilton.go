package graph

import (
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
func (g *Graph) Hamiltonian(startV int) []int {

	Am := copy2DSlice(g.Amatrix)

	path := make([]int, len(g.Amatrix))

	// Инициализируем путь как -1
	for i := range path {
		path[i] = -1
	}

	// Начинаем с вершины 0
	path[0] = startV
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

func (g *Graph) GamiltonTransform() {

	log.Print("func (G *Graph) EilerTransform() ")
	log.Print("делаем граф неориентированным")
	g.OrientToUnoriet()

	/*
		пока index, минимальный элемент из списка степеней вершин < n/2 {

			добавить столько вершин, чтобы его степень была n/2
		}
	*/

	{
		n := len(g.Amatrix)
		for v := 0; v < n; v++ {
			for i := 0; i < n; i++ {
				if g.degree(v) >= n/2 {
					break
				}
				if v != i && g.Amatrix[v][i] == 0 {
					g.Amatrix[v][i] = 1
					g.Amatrix[i][v] = 1
				}
			}
		}
	}

}

func (g *Graph) HamiltonWeight() (w int) {
	path := g.Hamiltonian(0)

	for i := 1; i < len(path); i++ {
		w += g.Amatrix[path[i-1]][path[i]]
		//fmt.Printf("%d += weight(%d, %d)\n", w, path[i-1], path[i])
	}

	return w
}
