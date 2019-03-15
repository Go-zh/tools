// The lostcancel command applies the github.com/Go-zh/tools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/Go-zh/tools/go/analysis/passes/lostcancel"
	"github.com/Go-zh/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
