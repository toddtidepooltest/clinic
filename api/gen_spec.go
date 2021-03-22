// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//+lraytEjoAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

