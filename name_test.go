package name

import (
	"net/http"
	"testing"

	"github.com/matryer/is"
	"github.com/nogoegst/name/.pkg-test"
)

type SimpleStruct struct {
}

func simpleFunc() {
}

var x int = 42

func TestNameOf(t *testing.T) {
	is := is.New(t)
	is.Equal(Of(SimpleStruct{}), "SimpleStruct")
	is.Equal(Of(simpleFunc), "simpleFunc")
	is.Equal(Of(x), "")
	is.Equal(Of(http.HandlerFunc(pkgtest.HandlerFunction)), "HandlerFunction")
	is.Equal(Of(http.HandlerFunc((&pkgtest.HandlerStruct{}).HandlerMethod)), "HandlerMethod")
	is.Equal(Of(&pkgtest.HandlerStruct{}), "HandlerStruct")
	is.Equal(Of(pkgtest.HandlerStructDirect{}), "HandlerStructDirect")
}
