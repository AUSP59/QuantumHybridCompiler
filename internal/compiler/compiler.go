// SPDX-License-Identifier: Apache-2.0
package compiler

import "qhcompiler/internal/translator"

type QuantumOp struct {
	Gate   string
	Qubits []string
}

func Compile(ir []translator.IRInstruction) []QuantumOp {
	var qops []QuantumOp
	for _, instr := range ir {
		switch instr.Opcode {
		case "GATE":
			qops = append(qops, QuantumOp{
				Gate:   instr.Params[0],
				Qubits: instr.Params[1:],
			})
		}
	}
	return qops
}
