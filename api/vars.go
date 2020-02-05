package main

import (
	"net/http"
	"sync"
)

var (
	varLock sync.RWMutex
	vars    map[*http.Request]map[string]interface{}
)

func OpenVars(r *http.Request) {
	varLock.Lock()
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varLock.RUnlock()
}

func main() {

}
