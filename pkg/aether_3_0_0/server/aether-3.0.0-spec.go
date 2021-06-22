// Code generated by oapi-codegen. DO NOT EDIT.
// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.0 DO NOT EDIT.
package server

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
	"ZFlxHNhgfu/x8f8DAAD//0qruI+StwAA",
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
