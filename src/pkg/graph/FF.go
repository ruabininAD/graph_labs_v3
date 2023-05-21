package graph

import (
	"fmt"
	"log"
)

func (G *Graph) FordFalkerson() [][]int {

	if G.Flags["Network"] == false {
		log.Print("граф не является сетью")
		fmt.Println("граф не является сетью")
		return nil
	}

	if G.Flags["BandwidthMatrix"] == false {
		log.Print("Матрица пропускной способности не задана")
		fmt.Println("Матрица пропускной способности не задана")
		return nil
	}
	// Функция для поиска матрицы максимального потока в графе

	source := 0
	sink := G.vCount - 1 // может vCount-1

	Am := G.Amatrix
	Bm := G.BandwidthMatrix

	n := len(Am)
	flow_matrix := make([][]int, n) // Матрица максимального потока
	for i := range flow_matrix {
		flow_matrix[i] = make([]int, n)
	}
	residual_matrix := make([][]int, n) // Остаточная сеть
	for i := range residual_matrix {
		residual_matrix[i] = make([]int, n)
		for j := range residual_matrix[i] {
			residual_matrix[i][j] = Bm[i][j]
		}
	}

	for {
		// Находим увеличивающий путь в остаточной сети с помощью алгоритма BFS
		parent := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		parent[source] = -2
		queue := []int{source}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			for nextNode, capacity := range residual_matrix[node] {
				if capacity > 0 && parent[nextNode] == -1 {
					parent[nextNode] = node
					queue = append(queue, nextNode)
				}
			}
		}

		// Если больше нет увеличивающих путей, заканчиваем работу
		if parent[sink] == -1 {
			break
		}

		// Находим минимальную пропускную способность на увеличивающем пути
		min_capacity := 1_000_000_000 // Большое число
		node := sink
		for node != source {
			parent_node := parent[node]
			min_capacity = min2(min_capacity, residual_matrix[parent_node][node])
			node = parent_node
		}

		// Увеличиваем поток вдоль увеличивающего пути и обновляем остаточную сеть
		node = sink
		for node != source {
			parent_node := parent[node]
			flow_matrix[parent_node][node] += min_capacity
			//flow_matrix[node][parent_node] -= min_capacity // для не ориентированных ребер снижаем в обе стороны
			residual_matrix[parent_node][node] -= min_capacity
			residual_matrix[node][parent_node] += min_capacity
			node = parent_node
		}
	}

	return flow_matrix
	//return makeUpperTriangular(flow_matrix)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (G *Graph) CostMaxFlow() (CostFlow int, Flow int) {
	//возвращает стоимость максимального потока

	if G.Flags["Network"] == false {
		fmt.Println("Граф не является сетью. Ошибка")
		return -1, -1
	}

	//сохраняем матрицы
	Am := copy2DSlice(G.Amatrix)
	Bm := copy2DSlice(G.BandwidthMatrix)

	log.Printf("CostMaxFlow()")

	log.Printf("матрица пропускных способностей:\n")
	for _, row := range G.BandwidthMatrix {
		log.Printf("%v \n", row)
	}
	log.Printf("матрица смежности:\n")
	for _, row := range G.Amatrix {
		log.Printf("%v \n", row)
	}

	for true {

		dlen, path, _ := G.BellmanFord(0, G.vCount-1)
		if len(path) < 2 {
			break
		}

		//находим пропускную способность дуг, по которым мы пошли
		vBmatris := make([]int, 0)
		for i := 1; i < len(path); i++ {
			//i-1, i
			a := path[i-1]
			b := path[i]
			vBmatris = append(vBmatris, G.BandwidthMatrix[a][b])
		}

		//находим самое узкое место.
		minvBmatris := 10000
		for _, v := range vBmatris {
			if minvBmatris > v {
				minvBmatris = v
			}
		}

		//уменьшаем значение в матрице пропускных способностей
		for i := 1; i < len(path); i++ {
			a := path[i-1]
			b := path[i]
			G.BandwidthMatrix[a][b] -= minvBmatris
			if G.BandwidthMatrix[a][b] == 0 {
				G.Amatrix[a][b] = 0
			}
		}

		// добавляем к новой стоимости стоимость прохода для 1 единицы товара * чило единиц товара
		CostFlow += dlen * minvBmatris
		Flow += minvBmatris

		//логирование по итерациям
		log.Printf("dlen: %v \n", dlen)
		log.Printf("path: %v \n", path)
		log.Printf("пропускная способность дуг: %v \n", vBmatris)
		log.Printf("самое узкое место: %v \n", minvBmatris)
		/*log.Printf("новая матрица пропускных способностей:\n")
		for _, row := range G.BandwidthMatrix {
			log.Printf("%v \n", row)
		}
		log.Printf("новая матрица смежности:\n")
		for _, row := range G.Amatrix {
			log.Printf("%v \n", row)
		}*/
		log.Printf("res: %d \n", CostFlow)
		log.Printf("Flow: %d \n", Flow)
		log.Printf(" ")

	}

	//восстанавливаем матрицы
	G.Amatrix = Am
	G.BandwidthMatrix = Bm
	return CostFlow, Flow
}

func (G *Graph) MinCostFlow(myFlow int) int {

	//Вычислить поток минимальной стоимости (в качестве величины потока брать значение, равное [2/3*max], где max – максимальный поток)
	//возвращает стоимость максимального потока
	if G.Flags["Network"] == false {
		fmt.Println("Граф не является сетью. Ошибка")
		return -1
	}

	resCost := 0
	resFlow := 0
	Am := copy2DSlice(G.Amatrix)
	Bm := copy2DSlice(G.BandwidthMatrix)

	log.Printf("MinCostFlow()")

	/*log.Printf("матрица пропускных способностей:\n")
	for _, row := range G.BandwidthMatrix {
		log.Printf("%v \n", row)
	}
	log.Printf("матрица смежности:\n")
	for _, row := range G.Amatrix {
		log.Printf("%v \n", row)
	}*/

	for true {

		dlen, path, _ := G.BellmanFord(0, G.vCount-1)
		if len(path) < 2 {
			break
		}

		//находим пропускную способность дуг, по которым мы пошли
		vBmatris := make([]int, 0)
		for i := 1; i < len(path); i++ {
			//i-1, i
			a := path[i-1]
			b := path[i]
			vBmatris = append(vBmatris, G.BandwidthMatrix[a][b])
		}

		//находим самое узкое место.
		minvBmatris := int(^uint(0) >> 1)
		for _, v := range vBmatris {
			if minvBmatris > v {
				minvBmatris = v
			}
		}

		//уменьшаем значение в матрице пропускных способностей
		for i := 1; i < len(path); i++ {
			a := path[i-1]
			b := path[i]
			G.BandwidthMatrix[a][b] -= minvBmatris
			if G.BandwidthMatrix[a][b] == 0 {
				G.Amatrix[a][b] = 0
			}
		}

		// добавляем к итоновой стоимости стоимость прохода для 1 единицы товара * чило единиц товара
		//log.Printf("resFlow+minvBmatris < myFlow === %d + %d < %d  === %d \n", resFlow, minvBmatris, myFlow, resFlow+minvBmatris < myFlow)
		if resFlow+minvBmatris < myFlow {
			resFlow += minvBmatris
			resCost += dlen * minvBmatris
		} else {
			resCost += dlen * (myFlow - resFlow)
			resFlow = myFlow
			log.Printf("end resCost: %d \n", resCost)
			G.Amatrix = Am
			G.BandwidthMatrix = Bm
			return resCost
			break
		}

		//log.Printf("dlen: %v \n", dlen)
		//log.Printf("path: %v \n", path)
		//log.Printf("пропускная способность дуг: %v \n", vBmatris)
		//log.Printf("самое узкое место: %v \n", minvBmatris)

		/*log.Printf("новая матрица пропускных способностей:\n")
		for _, row := range G.BandwidthMatrix {
			log.Printf("%v \n", row)
		}
		log.Printf("новая матрица смежности:\n")
		for _, row := range G.Amatrix {
			log.Printf("%v \n", row)
		}*/

		/*log.Printf("res: %d \n", resCost)
		log.Printf("Flow: %d \n", resFlow)
		log.Printf(" ")*/
	}

	return resCost
}

func copy2DSlice(s [][]int) [][]int {
	n := len(s)
	newS := make([][]int, n)
	for i := 0; i < n; i++ {
		m := len(s[i])
		newS[i] = make([]int, m)
		for j := 0; j < m; j++ {
			newS[i][j] = s[i][j]
		}
	}
	return newS
}
