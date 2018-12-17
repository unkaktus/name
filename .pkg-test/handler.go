// handler.go - this package deliberately has a dot and a hypen
// in it name to test name parsing.

package pkgtest

import (
	"net/http"
)

func HandlerFunction(w http.ResponseWriter, r *http.Request) {
}

type HandlerStruct struct {
}

func (h *HandlerStruct) HandlerMethod(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type HandlerStructDirect struct {
}

func (h HandlerStructDirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
