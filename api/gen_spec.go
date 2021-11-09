// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"K1tcLf43AAD//wyD+2gcYgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

