package main

const (
	TSTypeNumber string = "number"
	TSTypeString string = "string"
	TSTypeBool   string = "boolean"
	TSTypeAny    string = "any"
)

func typeMappings() map[string]string {
	return map[string]string{
		"int":        TSTypeNumber,
		"int8":       TSTypeNumber,
		"int16":      TSTypeNumber,
		"int32":      TSTypeNumber,
		"int64":      TSTypeNumber,
		"unit":       TSTypeNumber,
		"unit8":      TSTypeNumber,
		"unit16":     TSTypeNumber,
		"unit32":     TSTypeNumber,
		"unit64":     TSTypeNumber,
		"float32":    TSTypeNumber,
		"float64":    TSTypeNumber,
		"complex64":  TSTypeNumber,
		"complex128": TSTypeNumber,
		"string":     TSTypeString,
		"bool":       TSTypeBool,
		"byte":       TSTypeNumber,
		"rune":       TSTypeNumber,
	}
}

func GetIdentMapping(typeName string) string {
	mappings := typeMappings()
	if val, ok := mappings[typeName]; ok {
		return val
	}
	return TSTypeAny
}
