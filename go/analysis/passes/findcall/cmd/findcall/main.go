// The findcall command runs the findcall analyzer.
package main

import (
	"github.com/Go-zh/tools/go/analysis/passes/findcall"
	"github.com/Go-zh/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(findcall.Analyzer) }
