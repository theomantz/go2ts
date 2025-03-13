package main

import (
	"fmt"
	"go/ast"
	"strings"
)

func processStruct(name string, structType *ast.StructType, exportedStructs map[string]bool) string {
	var tsFields []string
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			tsType := goTypeToTsType(field.Type, exportedStructs)
			jsonTag := getJSONTag(field.Tag)
			tsFieldName := jsonTag
			if tsFieldName == "" {
				tsFieldName = fieldName.Name
			}
			tsFields = append(tsFields, fmt.Sprintf("    %s: %s", tsFieldName, tsType))
		}
	}

	var outputString string
	if len(tsFields) > 0 {
		outputString += strings.Join(tsFields, ";\n") + ";"
	}
	return fmt.Sprintf("export type %s = {\n%s\n};\n", name, outputString)
}

func getJSONTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	tagValue := tag.Value
	if strings.Contains(tagValue, "json:") {
		parts := strings.Split(tagValue, `"`)
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

func goTypeToTsType(expr ast.Expr, exportedStructs map[string]bool) string {
	switch t := expr.(type) {
	case *ast.Ident:
		if exportedStructs[t.Name] {
			return t.Name
		}
		return GetIdentMapping(t.Name)
	case *ast.ArrayType:
		return goTypeToTsType(t.Elt, exportedStructs) + "[]"
	case *ast.MapType:
		key := goTypeToTsType(t.Key, exportedStructs)
		value := goTypeToTsType(t.Value, exportedStructs)
		return fmt.Sprintf("{ [key: %s]: %s }", key, value)
	case *ast.StructType:
		// Inline anonymous struct
		var tsFields []string
		for _, field := range t.Fields.List {
			for _, fieldName := range field.Names {
				tsType := goTypeToTsType(field.Type, exportedStructs)
				tsFields = append(tsFields, fmt.Sprintf("    %s: %s", fieldName.Name, tsType))
			}
		}
		return fmt.Sprintf("{\n%s\n}", strings.Join(tsFields, "\n"))
	default:
		return "any" // Fallback for unsupported types
	}
}
