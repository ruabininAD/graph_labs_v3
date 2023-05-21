package graph

import (
	"fmt"
	"log"
	"math"
)

func (G *Graph) ShimbelStep(step int, fun string) *Graph {
	//step соответствует степени
	res := G
	for i := 0; i < step-1; i++ {
		res = ShimbelMultiply(res, G, fun)
	}
	return res
}

func (G *Graph) ShimbelDistanceMatrix(fun string) *Graph {
	log.Printf("start ShimbelDistanceMatrix  fun =%s", fun)
	// min  минимальный маршрут от точки до точки
	// max  максимальный маршрут от точки до точки
	res, err := NewGraph(G.GetVCount())
	res.Flags["oriented"] = true
	if err != nil {
		log.Print(err)
	}

	ShimbelSteps := make([]*Graph, 0)
	for i := 1; i < G.GetVCount(); i++ {
		tmp := G.ShimbelStep(i, fun)
		ShimbelSteps = append(ShimbelSteps, tmp)
	}

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			arrIJ := make([]int, G.vCount)
			log.Printf("start generated tab i = %v, j =%v", i, j)
			for step := 0; step < G.vCount-1; step++ {
				//log.Printf("append(arrIJ, ShimbelSteps[step].Get(i, j)) i = %v, j =%v", i, j)
				arrIJ = append(arrIJ, ShimbelSteps[step].Get(i, j))
			}

			value := 0
			if fun == "max" {
				value = max(arrIJ)
			} else {
				value = min(arrIJ)
			}
			log.Printf("%d, %d : %d", i, j, value)
			res.Set(i, j, value)
		}
	}

	return res
}

func ShimbelMultiply(b, a *Graph, fun string) *Graph {
	if a.vCount != b.vCount {
		log.Printf("The matrices cannot be multiplied: n1 =%v, n2 =  %v", a.vCount, b.vCount)
		panic("The matrices cannot be multiplied")
	}

	result, err := NewGraph(a.vCount)

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < a.vCount; i++ {
		for j := 0; j < a.vCount; j++ {
			arr := make([]int, a.vCount) //тут массив
			for k := 0; k < a.vCount; k++ {
				if a.Amatrix[i][k] == 0 || b.Amatrix[k][j] == 0 {
					arr = append(arr, 0)
					continue
				}
				arr = append(arr, a.Amatrix[i][k]+b.Amatrix[k][j])
				// добавление в массив
			}
			res := 0
			if fun == "max" {
				res = max(arr)
			} else {
				res = min(arr)
			}

			result.Amatrix[i][j] = res // выбор наибольшего из массива.
		}
	}

	return result
}

func (G *Graph) CountPaths(start int, end int) (count int, paths []int, weight int) {

	shortestPath := math.MaxInt32
	n := G.vCount

	var dfs func(current int, path []int)
	dfs = func(current int, path []int) {
		if current == end {
			count++
			shortestPath = minAB(shortestPath, len(path)-1)
			paths = path
			return
		}
		for i := 0; i < n; i++ {
			if G.Amatrix[current][i] != 0 && !contains(path, i) {
				dfs(i, append(path, i))
			}
		}
	}

	dfs(start, []int{start})
	if count == 0 {
		return -1, nil, -1
	}

	for i, _ := range paths {
		if i == 0 {
			continue
		}
		weight += G.Get(paths[i-1], paths[i])
	}

	return count, paths, weight
}

func contains(lst []int, el int) bool {
	for _, v := range lst {
		if v == el {
			return true
		}
	}
	return false
}

func minAB(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countNZero(arrRoad []int) (countNZero int) {
	count := 0
	log.Printf("проверка на нули массива: %q\n", fmt.Sprint(arrRoad))
	for lenRoad, countRoad := range arrRoad {
		if lenRoad != 0 {
			log.Printf("имеется %v дорог длинной в %v\n", countRoad, lenRoad)
			count += 1
		}
	}
	return count
}
