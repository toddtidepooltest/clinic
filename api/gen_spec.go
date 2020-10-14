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

