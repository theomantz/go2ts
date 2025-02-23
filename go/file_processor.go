package main

import (
	"go/ast"
	"go/token"
	"strings"
)

func processFile(file *ast.File) map[string]string {
	tsDefs := make(map[string]string)
	exportedStructs := make(map[string]bool)
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						if shouldExport(genDecl.Doc) {
							exportedStructs[typeSpec.Name.Name] = true
						}
					}
				}
			}
		}
	}
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						if shouldExport(genDecl.Doc) {
							tsDefs[typeSpec.Name.Name] = processStruct(typeSpec.Name.Name, structType, exportedStructs)
						}
					}
				}
			}
		}
	}
	return tsDefs
}

func shouldExport(doc *ast.CommentGroup) bool {
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
