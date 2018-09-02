package handler

import (
	"net/http"

	"github.com/alexandrevilain/monit/config"
	"github.com/alexandrevilain/monit/job"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type handler struct {
	store     job.Store
	responses *config.HttpConfig
}

// New creates a new instance of the handler struct
func New(store job.Store, responses *config.HttpConfig) http.Handler {
	h := handler{store, responses}
	r := mux.NewRouter()
	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/services/{name}", h.getServiceStatus).Methods("GET")
	return r
}

func (h *handler) getServiceStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, err := h.store.GetState(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unable to find this service"))
		return
	}
	content := h.responses.ErrorResponse
	if status {
		content = h.responses.WorkingResponse
	}
	w.Write([]byte(content))
}
