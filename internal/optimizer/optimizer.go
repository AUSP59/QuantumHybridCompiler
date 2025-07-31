// SPDX-License-Identifier: Apache-2.0
package optimizer

import "qhcompiler/internal/compiler"

func Optimize(qops []compiler.QuantumOp) []compiler.QuantumOp {
	var optimized []compiler.QuantumOp
	for i, op := range qops {
		if i > 0 && qops[i-1].Gate == op.Gate && equalQubits(qops[i-1].Qubits, op.Qubits) {
			continue // remove redundant consecutive
		}
		optimized = append(optimized, op)
	}
	return optimized
}

func equalQubits(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
