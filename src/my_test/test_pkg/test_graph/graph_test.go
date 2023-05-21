package test_graph

//go test -bench -m . .\graph_test.go

import (
	"graph_labs/src/pkg/generator"
	"graph_labs/src/pkg/graph"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestNewGraph(t *testing.T) { // fixme

	// arrange
	input := 1

	expected, err := graph.NewGraph(input)
	if err != nil {
		t.Error(err)
	}

	// act
	res, err := graph.NewGraph(input)
	if err != nil {
		t.Error(err)
	}

	// asserts

	if res.Amatrix[0][0] != expected.Amatrix[0][0] {
		t.Errorf("не корректное сравнение графов")
	}

}

// Бенчмарк для функции myFunction go test -bench . .\graph_test.go
func BenchmarkBellmanFord(b *testing.B) {
	logFile, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	Graph, _ := generator.NewErlingAcyclicOrientedGraph(10)
	rint := rand.New(rand.NewSource(time.Now().UnixNano())).Intn

	for i := 0; i < b.N; i++ {
		// Вызов функции, которую необходимо измерять
		result1, result2, _ := Graph.BellmanFord(rint(10), rint(10))
		// Используйте результат, чтобы предотвратить оптимизацию компилятора
		_, _ = result1, result2
	}
}

// Бенчмарк для функции myFunction
func BenchmarkDijcktra(b *testing.B) {
	logFile, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	Graph, _ := generator.NewErlingAcyclicOrientedGraph(10)
	rint := rand.New(rand.NewSource(time.Now().UnixNano())).Intn

	for i := 0; i < b.N; i++ {
		// Вызов функции, которую необходимо измерять
		result1, result2, err := Graph.Dijkstra(rint(10), rint(10))
		// Используйте результат, чтобы предотвратить оптимизацию компилятора
		_, _, _ = result1, result2, err
	}
}

func BenchmarkFloid(b *testing.B) {
	logFile, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
	Graph, _ := generator.NewErlingAcyclicOrientedGraph(10)
	for i := 0; i < b.N; i++ {
		// Вызов функции, которую необходимо измерять
		result1, result2, myMap := Graph.Floid()
		// Используйте результат, чтобы предотвратить оптимизацию компилятора
		_, _, _ = result1, result2, myMap
	}

}

func BenchmarkPrim(b *testing.B) {
	logFile, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
	Graph, _ := generator.NewErlingNetwork(10)
	for i := 0; i < b.N; i++ {
		// Вызов функции, которую необходимо измерять
		result1 := Graph.PrimAlgorithm()
		// Используйте результат, чтобы предотвратить оптимизацию компилятора
		_ = result1
	}

}

func BenchmarkKraskala(b *testing.B) {
	logFile, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
	Graph, _ := generator.NewErlingNetwork(10)
	for i := 0; i < b.N; i++ {
		// Вызов функции, которую необходимо измерять
		result1 := Graph.KruskalAlgorithm()
		// Используйте результат, чтобы предотвратить оптимизацию компилятора
		_ = result1
	}

}
