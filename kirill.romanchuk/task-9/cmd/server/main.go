package main

import (
	"os"

	"github.com/kirill.romanchuk/task-8/cmd/server/internal/config"
)

func main() {
	// parse config
	configFile, err := os.Open(*configPathFlag)
	if err != nil {
		panic(err)
	}

	defer configFile.Close()
	cfg, err := config.Unmarshal(configFile)
	if err != nil {
		panic(err)
	}

	// create rest server
	// run rest server
	// defer stop
}
