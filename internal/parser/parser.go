// SPDX-License-Identifier: Apache-2.0
package parser

import (
	"strings"
)

type Node struct {
	Command string
	Args    []string
}

func Parse(code string) []Node {
	lines := strings.Split(code, "\n")
	var nodes []Node

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) > 0 {
			nodes = append(nodes, Node{
				Command: parts[0],
				Args:    parts[1:],
			})
		}
	}

	return nodes
}
