package menu

import (
	"fmt"
	"graph_labs/src/pkg/generator"
	"graph_labs/src/pkg/graph"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
)

type Menu struct {
	logo  string
	graph *graph.Graph
}

func Cls() {
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout
	err := cls.Run()

	if err != nil {
		log.Printf("Ошибка очистки консоли\n")
	}
}

func (myMenu *Menu) MainMenu() {
	fmt.Printf("" +
		"0) дополнить\n" +
		"1) Сгенерировать новый граф\n" +
		"2) Показать граф\n" +
		"3) Число дорог из A в B\n" +
		"4) Показать матрицу Шимбала\n" +
		"5) Применить алгоритм Дейкстры\n" +
		"6) Применить алгоритм Беллмана-Форда\n" +
		"7) Применить алгоритм Флойда\n" +
		"8) Добавить веса\n" +
		"9) Алгоритм Форда фалкерсона\n" +
		"10) Стоимость максимального потока  \n" +
		"11) Поток минимальной стоимости \n" +
		"12) Остовное дерево минимальной стоимости\n" +
		"13) Найти число остовных подграфов\n" +
		"14) Закодировать остовное дерево минимальной стоимости кодом Прюфера\n" +
		"15) Лаба 6\n")
	choiceMainMenu := 0
	_, _ = fmt.Scan(&choiceMainMenu)
	Cls()
	switch choiceMainMenu {
	case 0:
		myMenu.graph.PrintLabel("граф: ")

		fmt.Println("x y v b")
		var x, y, v, b int

		_, _ = fmt.Scan(&x, &y, &v, &b)

		myMenu.graph.Amatrix[y][x] = v
		myMenu.graph.BandwidthMatrix[y][x] = b

	case 1:
		myMenu.Generated()
	case 2:
		myMenu.Print()
	case 3:
		a, b := 0, 0

		myMenu.graph.PrintLabel("граф:")

		fmt.Println("Введите начальную и конечную вершины")
		_, _ = fmt.Scan(&a, &b)
		if a == b {
			fmt.Println("a == b")
			return
		}

		if a > myMenu.graph.GetVCount() || b > myMenu.graph.GetVCount() {
			fmt.Println("Bершины с таким индексом нет")
			break
		}

		count, path, weight := myMenu.graph.CountPaths(a, b)
		if count == -1 {

			fmt.Printf("из %d в %d нет пути.\n", a, b)

		} else {

			fmt.Printf("из %d в %d есть %d путей. Минимальный путь:  %v  вес: %d \n", a, b, count, path, weight)

		}
	case 4:
		fmt.Printf("" +
			"1) результирующая матрица шимбала\n" +
			"2) шаг матрицы шимбала\n")
		choiseShimbal := 0
		_, _ = fmt.Scan(&choiseShimbal)
		switch choiseShimbal {
		case 1:
			fmt.Println("введите функцию min или max\n")
			fun := ""
			_, _ = fmt.Scan(&fun)
			myMenu.graph.ShimbelDistanceMatrix(fun).PrintLabel("Результурующая матрица шимбала для функции " + fun)

		case 2:
			fmt.Println("введите шаг для матрицы Шимбала\n")
			choiseStep := 0
			_, _ = fmt.Scan(&choiseStep)
			fmt.Println("введите функцию min или max\n")
			fun := ""
			_, _ = fmt.Scan(&fun)
			myMenu.graph.ShimbelStep(choiseStep, fun).PrintLabel("Матрица Шимбала для шага " + strconv.Itoa(choiseStep) + "и функции " + fun)
		}
	case 5:

		myMenu.graph.PrintLabel("граф:")

		if myMenu.graph.Flags["negativeWeight"] == true {
			fmt.Println("Невозможно применить алгоритм дейкстры для графа с отрицательными весами\n")
			break
		}
		fmt.Println("Введите стартовую вершину для алгоритма Дейкстры\n")
		startV := 0
		_, _ = fmt.Scan(&startV)

		fmt.Println("Введите конечную вершину для алгоритма Дейкстры\n")
		finishV := 0

		_, _ = fmt.Scan(&finishV)

		if startV > myMenu.graph.GetVCount() || finishV > myMenu.graph.GetVCount() {

			fmt.Println("нет такой вершины")
			break

		}

		distance, path, err := myMenu.graph.Dijkstra(startV, finishV)
		if err != nil {
			log.Print(err)
			fmt.Println("ошибка")
		}

		if len(path) == 0 {
			fmt.Printf("между вершинами %d и  %d пути нет\n", startV, finishV)
			break
		}

		fmt.Printf("между вершинами %d и  %d путь длинной %d: %v ", startV, finishV, distance, path)
	case 6:

		myMenu.graph.PrintLabel("граф:")

		fmt.Println("Введите стартовую вершину для алгоритма Беллмана Форда\n")
		startV := 0
		_, _ = fmt.Scan(&startV)

		fmt.Println("Введите конечную вершину для алгоритма  Беллмана Форда\n")
		finishV := 0
		_, _ = fmt.Scan(&finishV)

		if startV > myMenu.graph.GetVCount() {
			fmt.Println("нет такой вершины")
			break
		}
		path, distance, err := myMenu.graph.BellmanFord(startV, finishV)
		if err != nil {
			log.Print(err)
			fmt.Println("ошибка")
		}

		fmt.Printf("между вершинами %d и  %d путь длинной %d: %v ", startV, finishV, distance, path)
	case 7:

		myMenu.graph.PrintLabel("граф:")

		dist, next, paths := myMenu.graph.Floid()
		printResFloid(dist, next, paths)

	case 8:
		fmt.Printf("" +
			"1) только положительные значения\n" +
			"2) любые значения\n" +
			"3) Сделать веса единичными\n" +
			"4) Сгенерировать матрицу пропускных способностей\n")
		chioceWeight := 0
		_, _ = fmt.Scan(&chioceWeight)
		switch chioceWeight {
		case 1:
			myMenu.graph.SetRandomWeight("+")
		case 2:
			myMenu.graph.SetRandomWeight("-")
		case 3:

			myMenu.graph.RemoveWeights()
		case 4:
			myMenu.graph.SetRandomBandwidthMatrix(10)

		default:
			fmt.Println("не та кнопочка")
		}
	case 9:
		fmt.Println("Матрица пропускных способностей по алгоритму Форда Фалкерсона")
		myMenu.graph.PrintFordFalkerson()
	case 10:
		CostFlow, Flow := myMenu.graph.CostMaxFlow()
		if CostFlow == -1 && Flow == -1 {
			break
		}

		fmt.Printf("Максимальный поток %d \n", Flow)
		fmt.Printf("Стоимость максимального потока %d \n", CostFlow)
	case 11:
		_, Flow := myMenu.graph.CostMaxFlow()

		if Flow == -1 {
			break
		}

		fmt.Printf("Поток минимальной стоимости. \n"+
			"Поток = %d (2/3*max)\n"+
			"Стоимость потока = %d \n"+
			"", Flow*2/3, myMenu.graph.MinCostFlow(Flow*2/3))

	case 12:

		fmt.Printf("Остовное дерево минимальной стоимости по алгоритму Краскала\n")
		myMenu.graph.PrintKraskalaAlgorithm()

		fmt.Printf("Остовное дерево минимальной стоимости по алгоритму Прима\n")
		myMenu.graph.PrintPrimAlgorithm()
	case 13:
		myMenu.graph.Kirchhoff_Main()
	case 14:
		fmt.Printf("Остовное дерево минимальной стоимости по алгоритму Краскала\n")
		myMenu.graph.PrintKraskalaAlgorithm()
		PruferCode := myMenu.graph.PruferCode()
		fmt.Printf("Код прюфера: %v \n", PruferCode)
	case 15:
		if myMenu.graph.GetVCount() < 3 {
			fmt.Println("Слишком мало вершин ")
		} else {
			myMenu.lab6()
		}
	default:
		fmt.Println("не та кнопочка")
	}

	_, _ = fmt.Scanln()
}

func (myMenu *Menu) Generated() {
	fmt.Printf("Создать:\n" +
		"1)  ориентированный граф с помощью распределения Элдинга\n" +
		"2)  сеть с помощью распределения Элдинга\n" +
		"3)  декодировать код прюфера\n" +
		//"4) ориентированный ациклическй граф\n"
		"")

	graphVeriant := 0
	VCount := 0

	_, _ = fmt.Scan(&graphVeriant)
	Cls()
	fmt.Printf("количество вершин:\n")
	_, _ = fmt.Scan(&VCount)

	if TestVCount(VCount) == true {
	} else {
		return
	}

	var myGraph *graph.Graph
	var err error

	switch graphVeriant {
	case 1:

		myGraph, err = generator.NewErlingAcyclicOrientedGraph(VCount)
		if err != nil {
			fmt.Printf("Для %d вершин нельзя построить ациклический граф с %d ребер\n", VCount)
			return
		}
	case 2:
		myGraph, err = generator.NewErlingNetwork(VCount)
		if err != nil {
			fmt.Printf("Для %d вершин нельзя построить сеть с %d ребер\n", VCount)
			return
		}
	case 3:

		PrufferCode := make([]int, VCount-2)

		fmt.Println("ввод кода Прюфера:\n")
		a := 0
		for i := 0; i < VCount-2; i++ {
			_, _ = fmt.Scan(&a)
			PrufferCode[i] = a
		}
		fmt.Println(PrufferCode) //fixme
		myGraph, err = generator.PrufferGenerator(PrufferCode)
		if err != nil {
			fmt.Printf("Для %d вершин нельзя \n", VCount)
			return
		}
	default:
		myGraph, err = generator.NewErlingNetwork(1)
		fmt.Println("не та кнопка")
	}

	log.Printf("граф сгенерирован в консоли\n")

	myMenu.graph = myGraph

	myMenu.graph.PrintLabel("Граф сгенерирован:")

}

func (myMenu *Menu) lab6() {
	fmt.Printf("" +
		"1) проверить, является ли граф эйлеровым и гамильтоновым\n" +
		"2) модифицировать граф до эйлерова \n" +
		"3) модифицировать граф до гамильтонова\n" +
		"4) Построить эйлеров цикл\n" +
		"5) Построить гамильтонов цикл\n" +
		"6) Решить задачу коммивояжера на гамильтоновом графе\n")

	choiceMainMenu := 0
	_, _ = fmt.Scan(&choiceMainMenu)
	Cls()
	switch choiceMainMenu {
	case 1:
		myMenu.graph.PrintLabel("Граф:")

		myMenu.graph.OrientToUnoriet()
		// проверка на гамильтонов граф
		{
			HamiltonFlag := len(myMenu.graph.Hamiltonian(0)) >= 2

			fmt.Printf("Проверка на гамильтонов граф: %t\n", HamiltonFlag)
		}
		// проверка на эйлеров граф
		{
			EulerTour := myMenu.graph.EulerTour()

			EulerFlag := EulerTour[0] == EulerTour[len(EulerTour)-1]

			fmt.Printf("Проверка на эйлеров граф: %t\n", EulerFlag)
		}

		myMenu.graph.UnorientToOriet()
	//модифицировать граф до эйлерова
	case 2:
		myMenu.graph.EilerTransform()
		myMenu.graph.PrintLabel("сгенерированый эйлеров граф:")
	//модифицировать граф до гамильтонова
	case 3:
		myMenu.graph.GamiltonTransform()
		myMenu.graph.PrintLabel("сгенерированый гамильтонов граф:")
	//эйлеров цикл
	case 4:
		myMenu.graph.OrientToUnoriet()

		path := myMenu.graph.EulerTour()

		myMenu.graph.PrintLabel("Граф:")

		if len(path) < 3 {
			fmt.Println("в данном графе нет эйлерова пути")
		} else {
			fmt.Printf("эйлеров путь:  %v \n", path)
		}

		if path[0] == path[len(path)-1] {
			fmt.Printf("эйлеров путь является эйровым циклом:  %v \n", path)
		}
		fmt.Println()

		myMenu.graph.UnorientToOriet()
	//гамильтонов цикл
	case 5:
		myMenu.graph.OrientToUnoriet()

		path := myMenu.graph.Hamiltonian(0)

		myMenu.graph.PrintLabel("Граф:")

		if len(path) < 2 {
			fmt.Println("в данном графе нет гамильтонова цикла")
		} else {
			fmt.Printf("гамильтонов цикл:  %v", path)
		}

		fmt.Println()

		myMenu.graph.UnorientToOriet()

	case 6:
		myMenu.graph.OrientToUnoriet()

		path := myMenu.graph.Hamiltonian(0)

		myMenu.graph.PrintLabel("Граф:")

		if len(path) < 2 {

			fmt.Println("в данном графе нет гамильтонова цикла\n")
		} else {
			fmt.Printf("гамильтонов цикл:  %v\n", path)
			w := myMenu.graph.HamiltonWeight()

			fmt.Println("стоимость пути комивояжера ", w)
		}

		fmt.Println()

		myMenu.graph.UnorientToOriet()

	default:
		fmt.Println("не та кнопочка")
	}
}

func (myMenu *Menu) Print() {
	fmt.Println(
		"1) вывести в консоль\n" +
			"2) рендер\n" +
			"3) показать свойства\n" +
			"4) показать матрицу пропускных способностей\n")

	ChoisePrint := 0

	_, _ = fmt.Scan(&ChoisePrint)

	switch ChoisePrint {
	case 1:
		myMenu.graph.PrintLabel("граф:")

	case 2:
		myMenu.graph.Render()

	case 3:
		for key, v := range myMenu.graph.Flags {

			if v == true {
				fmt.Printf("%s ", key)
			}

		}
		fmt.Println()

	case 4:
		myMenu.graph.PrintLabelBandwidthMatrix("матрица пропускных способностей:")

	default:
		fmt.Println("не та кнопка")
	}
}

func ConsoleMenu() {
	Cls()
	var menu Menu
	for {
		Cls()
		menu.MainMenu()
		_, _ = fmt.Scanln()
	}

}

func printResFloid(dist, next [][]int, paths map[string][]int) {

	// вывод матрицы расстояний
	fmt.Println("Матрица расстояний:")

	for _, row := range dist {
		for _, v := range row {
			if v > math.MaxInt32-1000 {
				fmt.Printf("inf\t")
			} else {
				fmt.Printf("%d\t", v)
			}
		}
		fmt.Println()
	}

	// вывод матрицы следующих вершин на пути
	//fmt.Println("Матрица следующих вершин:")
	//for _, row := range next {
	//	fmt.Println(row)
	//}

	// вывод всех путей
	fmt.Println("Все пути:")
	for key, value := range paths {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func TestVCount(vc int) bool {
	if vc < 2 {
		return false
	} else {
		return true
	}

}
