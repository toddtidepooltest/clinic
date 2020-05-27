// Package Clinic provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AZa9bXNISreZbD7bzervAAAA//8bra78qi0AAA==",
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

