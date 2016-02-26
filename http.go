package example

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var upTime = time.Now().UTC().Unix()

func init() {
	r := httprouter.New()

	// routes
	r.GET("/service/status", ServiceStatus)

	http.Handle("/", r)
}
