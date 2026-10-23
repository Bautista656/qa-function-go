// Package hello is an HTTP function, registered by name for the Functions Framework.
package hello

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloHTTP", helloHTTP)
}

func helloHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from a Go function on CXF!")
}
