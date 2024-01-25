package main

import "awesomeProject/DoomParser/cmd/app/parser"

func main() {

	config, err := parser.ReadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	dependencies := parser.BuildDependencies(config)
	parser.RunApplication(dependencies)
}
