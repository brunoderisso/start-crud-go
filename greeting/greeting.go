package greeting

import (
	"fmt"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
