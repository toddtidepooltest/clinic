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
	"ZEU9KIpYmua0lNLN4n8CAAD//znJQnQvWgAA",
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

