package main

import (
	"graph_labs/src/interface/console/menu"
	"graph_labs/src/my_log"
	"os"
	//"graph_labs/src/pkg/graph"
)

func main() {
	my_log.SetLoger()

	arguments := os.Args[1:]
	if arguments[0] == "console" {
		menu.ConsoleMenu()
	}
}
