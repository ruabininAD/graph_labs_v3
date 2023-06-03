package graph

import (
	"fmt"
	"log"
)

// проверка граф на гамильтонов по Т Дирака и Т Оре
/*func (G *Graph) CheckGamilton() bool {

	log.Print("func (G *Graph) CheckGamilton() bool")
	log.Print("делаем граф неориентированным")
	G.OrientToUnoriet()
	defer G.UnorientToOriet()


	теорема Дирака
	FlagGamilton := dirac(G)

	теорема Оре
	flagOre := ore(G)

	return FlagGamilton && flagOre
}*/

// проверка граф на эйлеров по услов Достижимости и четности степени вершин
/*func (G *Graph) CheckEiler() bool {


		Граф является эйлеровым, если существует цикл,
		проходящий через каждое ребро графа ровно один раз.
		Этот цикл называется Эйлеровым циклом.

		Проверить, является ли граф эйлеровым, можно с помощью следующих критериев:

		Граф должен быть связным, то есть из любой вершины должно
		быть возможно достичь любой другой вершины по ребрам графа.

		Все вершины в графе должны иметь четную степень.
		Степень вершины — это количество ребер, которые с ней связаны.



			1) проверить с помошью матрицы шимбала. ее элементы должны быть != 0 кроме диагональных
		 	2) проверкой списка связности


	//Если граф удовлетворяет обоим условиям, то он является эйлеровым.
	log.Print("func (G *Graph) CheckEiler() bool ")
	log.Print("делаем граф неориентированным")
	G.OrientToUnoriet()
	defer G.UnorientToOriet()

	//Сохраняем матрицу смежности
	Am := copy2DSlice(G.Amatrix)

	Shm := copy2DSlice(G.ShimbelDistanceMatrix("min").Amatrix)

	//восстанавливаем матрицу смежности
	G.Amatrix = Am

	//флаг достижимости
	reachabilityFlag := true
	{
		for y := 0; y < G.vCount; y++ {
			for x := 0; x < G.vCount; x++ {
				if x != y && Shm[y][x] == 0 {

					reachabilityFlag = false

				}

			}
		}
	}

	//флаг четности
	parityFlag := true
	{
		for y := 0; y < G.vCount; y++ {

			if parityFlag == false {
				break
			}

			countC := 0
			for x := 0; x < G.vCount; x++ {

				if G.Amatrix[y][x] != 0 {
					countC += 1
				}

			}

			if countC%2 == 1 {
				parityFlag = false
				break
			}
		}
	}

	return reachabilityFlag && parityFlag
}*/

// Метод isEulerianCycle проверяет, есть ли Эйлеров цикл в графе.
// Он возвращает булево значение (есть ли цикл или нет) и стартовую вершину для цикла.
func (g *Graph) isEulerianCycle() (bool, int) {
	startVertex := 0
	oddCount := 0

	// Для каждой вершины в графе считаем степень вершины
	// (количество ребер, связанных с вершиной)
	for i, v := range g.Amatrix {
		degree := 0
		for _, w := range v {
			degree += w
		}
		// Если степень вершины нечетная, увеличиваем счетчик нечетных вершин
		// и обновляем стартовую вершину
		if degree%2 != 0 {
			oddCount++
			startVertex = i
		}
	}

	// Если количество вершин с нечетной степенью больше двух,
	// то Эйлеров цикл в графе отсутствует.
	if oddCount > 2 {
		return false, -1
	}

	return true, startVertex
}

// Метод EulerTour возвращает Эйлеров цикл или путь в графе.
func (g *Graph) EulerTour() []int {

	Am := copy2DSlice(g.Amatrix)

	ok, startVertex := g.isEulerianCycle()
	if !ok {
		//fmt.Println("Graph doesn't have a Eulerian Cycle.")
		g.Amatrix = (Am)
		return []int{}
	}

	// Инициализация стека и списка для хранения цикла
	var stack []int
	var cycle []int

	// Начинаем с выбранной стартовой вершины
	curV := startVertex
	// Пока стек не пуст или степень текущей вершины больше 0
	for len(stack) > 0 || g.degree(curV) > 0 {
		// Если степень текущей вершины равна 0
		// добавляем вершину в цикл и переходим к вершине на вершине стека
		if g.degree(curV) == 0 {
			cycle = append(cycle, curV)
			curV = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			// В противном случае добавляем вершину в стек
			// удаляем ребро и переходим к соседней вершине
			stack = append(stack, curV)
			neighbor := g.firstNeighbor(curV)
			g.Amatrix[curV][neighbor]--
			g.Amatrix[neighbor][curV]--
			curV = neighbor
		}
	}
	// Добавляем начальную вершину в конец цикла для замыкания
	cycle = append(cycle, curV)

	// Выводим Эйлеров цикл
	//fmt.Println("Eulerian Cycle: ", cycle)
	g.Amatrix = (Am)
	return cycle
}

func (g *Graph) EEulerTour() []int {

	adjMatrix := copy2DSlice(g.Amatrix)
	var cycle []int

	var dfs func(u int)
	dfs = func(u int) {
		for v := range adjMatrix[u] {
			if adjMatrix[u][v] > 0 {
				adjMatrix[u][v]--
				adjMatrix[v][u]--
				dfs(v)
			}
		}
		cycle = append(cycle, u)
	}

	for i, row := range adjMatrix {
		oddVertices := 0
		for _, val := range row {
			if val%2 == 1 {
				oddVertices++
			}
		}
		if oddVertices > 2 {
			return []int{} // нет эйлерова цикла
		}
		if oddVertices%2 == 1 {
			dfs(i) // начать с вершины нечетной степени, если такая есть
		}
	}
	if len(cycle) == 0 {
		dfs(0) // начать с вершины 0
	}

	return cycle
}

// Метод degree возвращает степень переданной вершины
func (g *Graph) degree(vertex int) int {
	degree := 0
	for _, w := range g.Amatrix[vertex] {
		degree += w
	}
	return degree
}

// Метод firstNeighbor возвращает первую соседнюю вершину для переданной вершины
func (g *Graph) firstNeighbor(vertex int) int {
	for i, w := range g.Amatrix[vertex] {
		if w > 0 {
			return i
		}
	}
	return -1
}

func (G *Graph) EilerTransform() {

	log.Print("func (G *Graph) EilerTransform() ")
	log.Print("делаем граф неориентированным")
	G.OrientToUnoriet()

	// соединяет все висячие вершины
	G.connectPendantVertices()

	//добавляет связи между вершинами с нечетной степенью
	G.addOrDelOdd(true)

	//удаляет связи между вершинами с нечетной степенью
	G.addOrDelOdd(false)

}

func listOddV(G *Graph) []int {
	oddV := make([]int, 0)
	{
		for y := 0; y < G.vCount; y++ {
			if G.degree(y)%2 == 1 {
				oddV = append(oddV, y)
			}
		}
	}
	return oddV
}

func getPairs(list []int, start int) [][]int {
	var pairs [][]int
	if start < len(list)-1 {
		for i := start + 1; i < len(list); i++ {
			pairs = append(pairs, []int{list[start], list[i]})
		}
		pairs = append(pairs, getPairs(list, start+1)...)
	}
	return pairs
}

// Функция возвращает список висячих вершин
func findPendantVertices(adjMatrix [][]int) []int {
	var pendantVertices []int

	for i, row := range adjMatrix {
		degree := 0
		for _, val := range row {
			degree += val
		}
		if degree == 1 {
			pendantVertices = append(pendantVertices, i)
		}
	}

	return pendantVertices
}

// Функция соединяет висячие вершины
func (G *Graph) connectPendantVertices() {
	PendantVertices := findPendantVertices(G.Amatrix)

	for len(PendantVertices) >= 2 {
		u, v := PendantVertices[0], PendantVertices[1]
		G.Amatrix[u][v] = 1
		G.Amatrix[v][u] = 1

		G.BandwidthMatrix[u][v] = 1
		G.BandwidthMatrix[v][u] = 1

		PendantVertices = PendantVertices[2:]
	}

	if len(PendantVertices) == 1 {
		u := PendantVertices[0]
		for v := range G.Amatrix {
			if u != v {
				G.Amatrix[u][v] = 1
				G.Amatrix[v][u] = 1
				break
			}
		}
	}
}

// Удаление и добавление связей для вершин с нечетной степенью.
// true - добавление flase-удаление
func (G *Graph) addOrDelOdd(val bool) {

	value := 0

	if val == true {
		value = 1
	}

	// список вершин с нечетной степенью
	oddV := listOddV(G)

	fmt.Printf("список вершин с нечетными степенями %d \n", oddV)
	G.PrintLabel("граф:")

	// все паросочетания нечетных вершин
	pairs := getPairs(oddV, 0)

	for _, pair := range pairs {
		if G.Amatrix[pair[0]][pair[1]] == (value+1)%2 {

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
				G.Amatrix[pair[0]][pair[1]] = value
				G.Amatrix[pair[1]][pair[0]] = value

				G.BandwidthMatrix[pair[0]][pair[1]] = value
				G.BandwidthMatrix[pair[1]][pair[0]] = value

				//удаляем вершины из списка не четных
				for i, v := range oddV {
					if v == pair[0] || v == pair[1] {
						oddV[i] = -1
					}

				}

				if val {
					fmt.Printf("добавленая связь %d - %d\n", pair[0], pair[1])
				} else {
					fmt.Printf("удалена связь %d - %d\n", pair[0], pair[1])
				}

			}

		}
	}
}
