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

