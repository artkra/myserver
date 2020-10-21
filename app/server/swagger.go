package server

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

const (
	pathField = "file"
	index     = "index.html"
)

// SwaggerDocHandler returns swagger.json
func SwaggerJSONHandler(
	w http.ResponseWriter, r *http.Request, pathParams map[string]string,
) {
	http.ServeFile(w, r, "./swagger/serve.swagger.json")
}

// SwaggerUIHandler(
func SwaggerUIHandler(
	w http.ResponseWriter, r *http.Request, pathParams map[string]string,
) {
	var filePath string
	if strings.ReplaceAll(r.URL.Path, "/", "") == "swagger" {
		filePath = index
	} else {
		filePath = path.Base(r.URL.Path)
	}
	http.ServeFile(w, r, fmt.Sprintf("./swagger/%s", filePath))
}
