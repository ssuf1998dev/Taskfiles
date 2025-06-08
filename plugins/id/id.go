package main

import (
	"github.com/extism/go-pdk"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

//go:wasmexport uuid
func _uuid() int32 {
	pdk.OutputString(uuid.New().String())
	return 0
}

//go:wasmexport uuidv6
func _uuidv6() int32 {
	v, err := uuid.NewV6()
	if err != nil {
		pdk.OutputString("")
	} else {
		pdk.OutputString(v.String())
	}
	return 0
}

//go:wasmexport uuidv7
func _uuidv7() int32 {
	v, err := uuid.NewV7()
	if err != nil {
		pdk.OutputString("")
	} else {
		pdk.OutputString(v.String())
	}
	return 0
}

type nanoidInput struct {
	Alphabet *string `json:"alphabet,omitempty"`
	Length   *int    `json:"length,omitempty"`
}

//go:wasmexport nanoid
func _nanoid() int32 {
	input := nanoidInput{}
	pdk.InputJSON(&input)

	lng := (func() int {
		if input.Length != nil {
			return *input.Length
		}
		return 21
	})()

	var v string
	var err error
	if input.Alphabet != nil {
		v, err = gonanoid.Generate(*input.Alphabet, lng)
	} else {
		v, err = gonanoid.New(lng)
	}

	if err != nil {
		pdk.OutputString("")
	} else {
		pdk.OutputString(v)
	}

	return 0
}
