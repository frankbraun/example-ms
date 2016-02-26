package example

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	MicroServiceName = "EXAMPLE"
)

func generateJsonResponse(w http.ResponseWriter, payload []byte, status int) {

	// write the Header Content-Type to json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// add the status Header
	w.WriteHeader(status)

	// write the marshalled payload
	w.Write(payload)

}

type serviceStatus struct {
	Message string `json:"message"`
	Uptime  string `json:"uptime"`
}

// API GET "/service/status"
// Expect a 200 StatusOK HTTP Response
func ServiceStatus(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// find duration for the upTime timestamp
	the_duration := time.Since(time.Unix(upTime, 0))

	// response
	st := &serviceStatus{
		Message: "Micro Service " + MicroServiceName + " Status is OK",
		Uptime:  the_duration.String(),
	}
	payload, _ := json.Marshal(st)
	generateJsonResponse(w, payload, http.StatusOK)
}
