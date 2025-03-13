package main

import (
	"go/ast"
	"strings"
)

func shouldExport(doc *ast.CommentGroup, marker string) bool {
	if doc == nil {
		return false
	}
	for _, comment := range doc.List {
		comment := strings.TrimSpace(comment.Text)
		if comment == "// @ts-export" || comment == "//@ts-export" {
			return true
		}
	}
	return false
}
