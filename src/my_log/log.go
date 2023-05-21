package my_log

import (
	"log"
	"os"
)

func SetLoger() {
	logFile, err := os.OpenFile("src/my_log/logfile.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)
}
