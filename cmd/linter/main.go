package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"linter/internal/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analizer)
}
