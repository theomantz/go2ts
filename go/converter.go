package main

import (
	"fmt"
	"go/ast"
	"strings"
)

func processStruct(name string, structType *ast.StructType, exportedStruct map[string]bool) string {
	var tsFields []string
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			tsType := goTypeToTsType(name, field.Type)
			jsonTag := getJSONTag(field.Tag)
			tsFieldName := jsonTag
			if tsFieldName == "" {
				tsFieldName = fieldName.Name
			}
			tsFields = append(tsFields, fmt.Sprintf("    %s: %s", tsFieldName, tsType))
		}
	}

	tsOutput := fmt.Sprintf("export type %s = {\n%s\n};\n", name, strings.Join(tsFields, ";\n"))
	return tsOutput
}

func getJSONTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	// Extract json tag value (e.g., `json:"id"`)
	tagValue := tag.Value
	if strings.Contains(tagValue, "json:") {
		parts := strings.Split(tagValue, `"`)
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

func goTypeToTsType(name string, expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return GetIdentMapping(t.Name)
	case *ast.ArrayType:
		return goTypeToTsType(name, t.Elt) + "[]"
	case *ast.MapType:
		key := goTypeToTsType(name, t.Key)
		value := goTypeToTsType(name, t.Value)
		return fmt.Sprintf("{ [key: %s]: %s }", key, value)
	case *ast.StructType:
		// TODO: support nested structs
	default:
		return "any" // Fallback for unsupported types
	}
}
