package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/firodj/perapian/common"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestAja(t *testing.T) {
	assert := assert.New(t)
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./samples/petstore.json")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, id := range doc.Paths.InMatchingOrder() {
		path := doc.Paths.Find(id)
		ops := path.Operations()
		for _, met := range common.HttpMethods {
			if _, ok := ops[met]; ok {
				key := fmt.Sprintf("%s%s", met, id)
				actual := strings.SplitN(key, "/", 2)

				assert.Equal(id, "/"+actual[1])
				assert.Equal(met, actual[0])
			}
		}

	}
}
