package name

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

func functionName(i interface{}) string {
	rawName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	l := strings.LastIndex(rawName, ".")
	r := strings.LastIndex(rawName, "-")
	if r < 0 || r < l {
		r = len(rawName)
	}
	return rawName[l+1 : r]
}

func structName(i interface{}) string {
	return reflect.Indirect(reflect.ValueOf(i)).Type().Name()
}

func handlerName(h http.Handler) string {
	if _, ok := h.(http.HandlerFunc); ok {
		return functionName(h)
	}
	return structName(h)
}

// Of returns the name of object i that can have it (struct, func).
// If i is an http.Handler, Of returns its meaningful name
// (not ServeHTTP).
func Of(i interface{}) string {
	if h, ok := i.(http.Handler); ok {
		return handlerName(h)
	}
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Func {
		return functionName(i)
	}
	if v.Kind() == reflect.Struct {
		return structName(i)
	}
	return ""
}
