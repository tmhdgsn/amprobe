package hook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/tmhdgsn/amprobe/alert"
	"github.com/tmhdgsn/amprobe/metrics"
)

type (
	Hook struct {
		sync.Mutex
		alerts []*alert.Message
		s      *http.Server
	}
)

func (h *Hook) alertsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getHandler(w, r)
	case http.MethodPost:
		h.postHandler(w, r)
	default:
		http.Error(w, "unsupported HTTP method", 400)
	}
}

func New(addr string) *Hook {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", addr),
	}
	return &Hook{s: server}
}

func (h *Hook) ListenAndServe(metricsHandler http.Handler) error {
	http.HandleFunc("/alerts", h.alertsHandler)
	http.Handle("/metrics", metricsHandler)
	return h.s.ListenAndServe()
}

func (h *Hook) getHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	h.Lock()
	defer h.Unlock()

	if err := enc.Encode(h.alerts); err != nil {
		log.Printf("error encoding messages: %v", err)
	}
}

func (h *Hook) postHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m alert.Message
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}

	h.Lock()
	defer h.Unlock()

	h.alerts = append(h.alerts, &m)
	log.Printf("received alert: %+v\n", m)
	metrics.AlertsProcessed.Inc()

}
