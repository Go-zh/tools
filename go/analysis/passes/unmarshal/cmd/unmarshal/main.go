// The unmarshal command runs the unmarshal analyzer.
package main

import (
	"github.com/Go-zh/tools/go/analysis/passes/unmarshal"
	"github.com/Go-zh/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unmarshal.Analyzer) }
