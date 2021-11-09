package main

import (
	"net/http"

	"github.com/tuckersGo/goWeb/tree/master/web5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
