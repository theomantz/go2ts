package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
)

var (
	dir = flag.String("dir", ".", "Input go directory")
	out = flag.String("out", "./go-types.ts", "Output file")
	marker = flag.String("marker", "// @ts-export", "Marker used to identify go structs to export")
)

func main() {

	flag.Usage = usage
	flag.Parse()
	// File or directory to parse
	inputDirPath := flag.Arg(0)
	outputFilePath := flag.Arg(1)
	marker := flag.Arg(2)
	fset := token.NewFileSet()

	// Parse the directory
	packages, err := parser.ParseDir(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Process each package
	for _, pkg := range packages {
		processPackage(pkg)
	}
}

func processPackage(fset *token.FileSet, pkg *ast.Package) string {
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
                            if shouldExport(genDecl.Doc) {
                                exportedStructs[typeSpec.Name.Name] = true
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
                            if shouldExport(genDecl.Doc) {
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
      ID   int    `json:"id"`
      Name string `json:"name"`
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
