package main

import "flag"

var (
	configPathFlag = flag.String("config", "config/default.yaml", "Path to YAML config")
)

func init() {
	flag.Parse()
}
