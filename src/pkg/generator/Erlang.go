package generator

import (
	"graph_labs/src/pkg/graph"
	"math"
	"math/rand"
	"time"

	"log"
)

func NewErlingAcyclicOrientedGraph(VCount int) (*graph.Graph, error) {
	log.Printf("func NewErlingGraph(%d int\n", VCount)

	G, _ := graph.NewGraph(VCount)

	G.Flags["acyclic"] = true
	G.Flags["oriented"] = true

	k := 5        // параметр формы распределения
	lambda := 2.0 // параметр интенсивности распределения
	res := 0.

	log.Printf("Распределение Элдинга для параметров\n "+
		"k = %v - параметр формы распределения \n "+
		"λ = %.2f -параметр интенсивности распределения \n", k, lambda)

	probabilities := make([]float64, 0)

	//формаирование вероятностей появления чисел от 0 до VCount
	for i := 0; i < VCount; i++ {
		res = erlangDistribution(float64(i), k, lambda)
		probabilities = append(probabilities, res)
	}
	log.Printf("Распределение вероятностей: %v", probabilities)
	//формирование ациклического графа на основе распределения вероятностей

	for i := 0; i < VCount; i++ {
		a := generateRandomNumber(probabilities)
		zeroSlise := make([]int, i+1)
		slise := createRandomSlice(VCount-i-1, a)

		slise = append(zeroSlise, slise...)
		for j, v := range slise {

			G.Set(i, j, v)

		}

	}
	return G, nil
}

func NewErlingNetwork(VCount int) (*graph.Graph, error) {

	G, err := NewErlingAcyclicOrientedGraph(VCount)
	if err != nil {
		log.Print("ошибка в построении NewErlingAcyclicOrientedGraph в NewErlingNetwork")
	}

	G.Flags["Network"] = true

	// проверка строки
	for rowIndex, row := range G.Amatrix {
		stockFlag := true //строка из нулей
		for _, val := range row {
			if val != 0 { //в строке есть не нули, тогда это не сток
				stockFlag = false // это уже не сток
				break
			}
		}

		if stockFlag {
			indexNewLine := rowIndex + rand.Intn(VCount-rowIndex)
			G.Amatrix[indexNewLine][rowIndex] = 1

		}
	}

	//проверка столбцов
	for x, _ := range G.Amatrix {
		sourseFlag := true // стоблец пустой ИСТОЧНИК
		for y, _ := range G.Amatrix {
			if G.Amatrix[y][x] != 0 {
				sourseFlag = false
				break
			}
		}
		if sourseFlag && x != 0 {
			yNewLine := rand.Intn(x)
			G.Amatrix[yNewLine][x] = 1

		}

	}

	G.Set(0, 0, 0)
	G.Set(VCount-1, VCount-1, 0)

	G.SetRandomBandwidthMatrix(10)
	return G, err

}

// Функция для вычисления плотности вероятности распределения Эрланга
func erlangDistribution(x float64, k int, lambda float64) float64 {
	numerator := math.Pow(lambda, float64(k)) * math.Pow(x, float64(k-1)) * math.Exp(-lambda*x)
	denominator := float64(factorial(k - 1))
	return numerator / denominator
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generateRandomNumber(probabilities []float64) int {
	sum := 0.0
	for _, p := range probabilities {
		sum += p
	}

	r := rand.Float64() * sum
	for i, p := range probabilities {
		if r < p {
			return i
		}
		r -= p
	}

	return len(probabilities) - 1
}

func createRandomSlice(n, m int) []int {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, n) // Создаем срез длины n
	if m > n {
		m = n
	}
	// Заполняем m случайных элементов в срезе
	for i := 0; i < m; i++ {
		index := random.Intn(n) // Генерируем случайный индекс
		if slice[index] == 1 {
			i--
			continue
		}
		slice[index] = 1 // Записываем значение в срез по случайному индексу
	}

	return slice
}
