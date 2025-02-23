Go2TS - Convert Go structs to TypeScript types

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
  type User = {
      id: number;
      name: string;
  };

Notes:
  - Only structs with the specified marker (e.g., "// @ts-export") are converted.
  - Nested structs marked with the same marker become separate TypeScript types.
  - Unexported nested structs are mapped to "any" by default.
  - Supported Go types: int, string, bool, arrays, slices, maps, etc.
  - JSON tags override field names in the output.

For more information, see: https://github.com/theomantz/go2ts
