// SPDX-License-Identifier: Apache-2.0
package main

import (
	"fmt"
	"os"

	"qhcompiler/internal/gui"
	"qhcompiler/internal/parser"
	"qhcompiler/internal/quantum"
	"qhcompiler/internal/compiler"
	"qhcompiler/internal/optimizer"
	"qhcompiler/internal/translator"
)

func main() {
	if len(os.Args) > 1 {
		sourceFile := os.Args[1]
		code, err := os.ReadFile(sourceFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		ast := parser.Parse(string(code))
		ir := translator.Translate(ast)
		qir := compiler.Compile(ir)
		opt := optimizer.Optimize(qir)
		sim := quantum.Emulate(opt)
		fmt.Println("Quantum Simulation Result:", sim)
	} else {
		gui.Launch()
	}
}
