package testdata

import "time"

// @ts-export
type ExportedTestStruct struct {
	// // Basic types
	// IntField        int        `json:"intField"`
	// Int8Field       int8       `json:"int8Field"`
	// Int16Field      int16      `json:"int16Field"`
	// Int32Field      int32      `json:"int32Field"`
	// Int64Field      int64      `json:"int64Field"`
	// UintField       uint       `json:"uintField"`
	// Uint8Field      uint8      `json:"uint8Field"`
	// Uint16Field     uint16     `json:"uint16Field"`
	// Uint32Field     uint32     `json:"uint32Field"`
	// Uint64Field     uint64     `json:"uint64Field"`
	// Float32Field    float32    `json:"float32Field"`
	// Float64Field    float64    `json:"float64Field"`
	// Complex64Field  complex64  `json:"complex64Field"`
	// Complex128Field complex128 `json:"complex128Field"`
	// StringField     string     `json:"stringField"`
	// BoolField       bool       `json:"boolField"`
	// ByteField       byte       `json:"byteField"` // alias for uint8
	// RuneField       rune       `json:"runeField"` // alias for int32

	// // Composite types
	// ArrayField  [3]int         `json:"arrayField"`
	// SliceField  []string       `json:"sliceField"`
	MapField map[string]int `json:"mapField"`
	// StructField struct {
	// 	Nested string `json:"nested"`
	// } `json:"structField"`

	NamedStructField NamedStructField

	// // Pointers
	// IntPtrField *int `json:"intPtrField"`

	// // Interfaces
	// AnyField interface{} `json:"anyField"`

	// // Custom type
	// CustomTypeField CustomType `json:"customTypeField"`

	// // Imported type
	// TimeField time.Time `json:"timeField"`

	// // Channel (not typically serialized, but included for completeness)
	// ChanField chan int `json:"chanField"`
}

type NamedStructField struct {
	SomeType   int
	SomeString string
}

// CustomType is a custom defined type for testing
type CustomType string

type UnexportedTestStruct struct {
	// Basic types
	IntField        int        `json:"intField"`
	Int8Field       int8       `json:"int8Field"`
	Int16Field      int16      `json:"int16Field"`
	Int32Field      int32      `json:"int32Field"`
	Int64Field      int64      `json:"int64Field"`
	UintField       uint       `json:"uintField"`
	Uint8Field      uint8      `json:"uint8Field"`
	Uint16Field     uint16     `json:"uint16Field"`
	Uint32Field     uint32     `json:"uint32Field"`
	Uint64Field     uint64     `json:"uint64Field"`
	Float32Field    float32    `json:"float32Field"`
	Float64Field    float64    `json:"float64Field"`
	Complex64Field  complex64  `json:"complex64Field"`
	Complex128Field complex128 `json:"complex128Field"`
	StringField     string     `json:"stringField"`
	BoolField       bool       `json:"boolField"`
	ByteField       byte       `json:"byteField"` // alias for uint8
	RuneField       rune       `json:"runeField"` // alias for int32

	// Composite types
	ArrayField  [3]int         `json:"arrayField"`
	SliceField  []string       `json:"sliceField"`
	MapField    map[string]int `json:"mapField"`
	StructField struct {
		Nested string `json:"nested"`
	} `json:"structField"`

	// Pointers
	IntPtrField *int `json:"intPtrField"`

	// Interfaces
	AnyField interface{} `json:"anyField"`

	// Custom type
	CustomTypeField CustomType `json:"customTypeField"`

	// Imported type
	TimeField time.Time `json:"timeField"`

	// Channel (not typically serialized, but included for completeness)
	ChanField chan int `json:"chanField"`
}
