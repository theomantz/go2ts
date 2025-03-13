package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestProcessPackage(t *testing.T) {
	tests := []struct {
		name     string
		input    string // Go source code
		expected string // Expected TS output
	}{
		{
			name: "Basic struct with primitives",
			input: `
                package test
                // @ts-export
                type Simple struct {
                    ID int ` + "`json:\"id\"`" + `
                    Name string ` + "`json:\"name\"`" + `
                }
            `,
			expected: `export type Simple = {
    id: number;
    name: string;
};
`,
		},
		{
			name: "Nested exported struct",
			input: `
                package test
                // @ts-export
                type Child struct {
                    Value string` + "`json:\"value\"`" + `
                }
                // @ts-export
                type Parent struct {
                    Child Child ` + "`json:\"child\"`" + `
                }
            `,
			expected: `export type Child = {
    value: string;
};

export type Parent = {
    child: Child;
};
`,
		},
		{
			name: "Nested unexported struct",
			input: `
                package test
                // @ts-export
                type Parent struct {
                    Child UnexportedChild ` + "`json:\"child\"`" + `
                }
                type UnexportedChild struct {
                    Value string ` + "`json:\"value\"`" + `
                }
            `,
			expected: `export type Parent = {
    child: any;
};
`,
		},
		{
			name: "Struct with composite types and pointer",
			input: `
                package test
                // @ts-export
                type Complex struct {
                    Numbers []int ` + "`json:\"numbers\"`" + `
                    Mapping map[string]bool ` + "`json:\"mapping\"`" + `
                    Ptr *string ` + "`json:\"ptr\"`" + `
                }
            `,
			expected: `export type Complex = {
    numbers: number[];
    mapping: { [key: string]: boolean };
    ptr: any;
};
`,
		},
		{
			name: "Empty struct",
			input: `
                package test
                // @ts-export
                type Empty struct {}
            `,
			expected: `export type Empty = {
};
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup file set and parse input
			fset := token.NewFileSet()
			file, err := parser.ParseFile(fset, "test.go", tt.input, parser.ParseComments)
			if err != nil {
				t.Fatalf("Failed to parse input: %v", err)
			}

			// Wrap file in a package
			pkg := &ast.Package{
				Name:  "test",
				Files: map[string]*ast.File{"test.go": file},
			}

			// Run the processing
			got := processPackage(fset, pkg, "// @ ts-export")

			// Normalize whitespace for comparison
			got = normalizeTS(got)
			expected := normalizeTS(tt.expected)

			if got != expected {
				t.Errorf("Test %q failed\nGot:\n%s\nExpected:\n%s", tt.name, got, expected)
			}
		})
	}
}

// normalizeTS removes extra whitespace and ensures consistent formatting for comparison
func normalizeTS(ts string) string {
	lines := strings.Split(strings.TrimSpace(ts), "\n")
	var cleaned []string
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			cleaned = append(cleaned, trimmed)
		}
	}
	return strings.Join(cleaned, "\n")
}
