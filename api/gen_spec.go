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

