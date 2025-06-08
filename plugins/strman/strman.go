package main

import (
	"github.com/extism/go-pdk"
	"github.com/gobeam/stringy"
)

//go:wasmexport camelCase
func camelCase() int32 {
	input := pdk.InputString()
	pdk.OutputString(stringy.New(input).CamelCase().Get())
	return 0
}

//go:wasmexport kebabCase
func kebabCase() int32 {
	input := pdk.InputString()
	pdk.OutputString(stringy.New(input).KebabCase().Get())
	return 0
}

//go:wasmexport snakeCase
func snakeCase() int32 {
	input := pdk.InputString()
	pdk.OutputString(stringy.New(input).SnakeCase().Get())
	return 0
}

//go:wasmexport PascalCase
func pascalCase() int32 {
	input := pdk.InputString()
	pdk.OutputString(stringy.New(input).PascalCase().Get())
	return 0
}
