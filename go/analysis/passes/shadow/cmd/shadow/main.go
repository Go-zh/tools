// The shadow command runs the shadow analyzer.
package main

import (
	"github.com/Go-zh/tools/go/analysis/passes/shadow"
	"github.com/Go-zh/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
