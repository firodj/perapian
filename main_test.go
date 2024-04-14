package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/getkin/kin-openapi/openapi3"
)

func TestAja(t *testing.T) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./samples/petstore.json")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	spew.Dump(doc.Info.License)

	t.Error("wew")
}
