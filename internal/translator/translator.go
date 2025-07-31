// SPDX-License-Identifier: Apache-2.0
package translator

import (
	"qhcompiler/internal/parser"
)

type IRInstruction struct {
	Opcode string
	Params []string
}

func Translate(nodes []parser.Node) []IRInstruction {
	var ir []IRInstruction
	for _, node := range nodes {
		switch node.Command {
		case "let":
			ir = append(ir, IRInstruction{Opcode: "LOAD", Params: node.Args})
		case "apply":
			ir = append(ir, IRInstruction{Opcode: "GATE", Params: node.Args})
		case "qif":
			ir = append(ir, IRInstruction{Opcode: "QIF", Params: node.Args})
		}
	}
	return ir
}
