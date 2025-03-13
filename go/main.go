package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	// File or directory to parse
	inputDirPath := flag.String("dir", ".", "directory containing Go source files")
	outputFilePath := flag.String("out", "types.ts", "output file for TypeScript definitions")
	marker := flag.String("marker", "// @ ts-export", "marker used to identify Go structs for export to TypeScript")
	fset := token.NewFileSet()

	// Parse the directory
	packages, err := parser.ParseDir(fset, *inputDirPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Process each package
	var tsOutput string
	for _, pkg := range packages {
		tsOutput += processPackage(fset, pkg, *marker)
	}

	err = os.WriteFile(*outputFilePath, []byte(tsOutput), 0644)
	if err != nil {
		log.Fatalf("failed to write to %s: %v", *outputFilePath, err)
	}

	log.Printf("definitions successfully written to %s", *outputFilePath)
}

func processPackage(fset *token.FileSet, pkg *ast.Package, marker string) string {
	// Map to store TypeScript definitions
	tsDefs := make(map[string]string)
	// Set of exported struct names
	exportedStructs := make(map[string]bool)

	// First pass: Identify all exported structs
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if _, ok := typeSpec.Type.(*ast.StructType); ok {
							if shouldExport(genDecl.Doc, marker) {
								exportedStructs[typeSpec.Name.Name] = true
								log.Printf("loop exportedStructs: %v", exportedStructs)
							}
						}
					}
				}
			}
		}
	}

	// Second pass: Process structs and generate TS definitions
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							if shouldExport(genDecl.Doc, marker) {
								log.Printf("exported exportedStructs: %v", exportedStructs)
								tsDefs[typeSpec.Name.Name] = processStruct(typeSpec.Name.Name, structType, exportedStructs)
							}
						}
					}
				}
			}
		}
	}

	// Combine all definitions into a single string
	var output strings.Builder
	for _, def := range tsDefs {
		output.WriteString(def)
		output.WriteString("\n")
	}
	return output.String()
}

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `Go2TS - Convert Go structs to TypeScript types

Usage:
  go2ts [flags]

Description:
  Go2TS scans a directory of Go source files, identifies structs marked with "// @ts-export",
  and generates corresponding TypeScript type definitions. It supports nested structs,
  basic Go types, and JSON tags for field naming.

Flags:
  -dir string
        Directory containing Go source files to process (default: ".")
  -out string
        Output file for TypeScript definitions (default: "types.ts")
  -marker string
        Comment marker to identify structs for export (default: "// @ts-export")
  -v    Enable verbose output

Examples:
  # Convert structs in the current directory, output to types.ts
  go2ts

  # Convert structs in ./src, output to custom-types.ts
  go2ts -dir ./src -out custom-types.ts

  # Use a custom marker and enable verbose logging
  go2ts -dir ./api -marker "// @typescript" -v

Input Example (Go):
  // File: ./example.go
  package example

  // @ts-export
  type User struct {
      ID   int    json:"id"
      Name string json:"name"
  }

Output Example (TypeScript):
  // File: types.ts
  export type User = {
      id: number;
      name: string;
  };

Notes:
  - Only structs with the specified marker (e.g., "// @ts-export") are converted.
  - Nested structs marked with the same marker become separate TypeScript types.
  - Unexported nested structs are mapped to "any" by default.
  - Supported Go types: int, string, bool, arrays, slices, maps, etc.
  - JSON tags override field names in the output.
  - TypeScript types are exported by default

For more information, see: https://github.com/theomantz/go2ts`
