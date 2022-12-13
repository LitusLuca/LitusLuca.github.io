//go:build wasm
// +build wasm

package loader

import (
	"io"
	"net/http"
	"syscall/js"
)



func ReadFile(path string) (io.Reader, error) {
	root := js.Global().Get("location").Get("origin").String()
	url := root + "/" + path
	res, err := http.Get(url)
	if err != nil {
		return	nil, err
	}
	return res.Body, nil
}
