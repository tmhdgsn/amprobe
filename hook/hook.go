package hook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tmhdgsn/amprobe/alert"
)

type (
	Hook struct {
		sync.Mutex
		alerts map[string]*AlertState
		s      *http.Server
	}

	AlertState struct {
		Received time.Time
		Msg      *alert.Message
	}
)

func (h *Hook) alertsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
	case http.MethodPost:
		h.post(w, r)
	case http.MethodHead:
		w.WriteHeader(http.StatusAccepted)
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

func (h *Hook) ListenAndServe() error {
	http.HandleFunc("/alerts", h.alertsHandler)
	http.Handle("/metrics", promhttp.Handler())
	return h.s.ListenAndServe()
}

func (h *Hook) get(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	h.Lock()
	defer h.Unlock()

	if err := enc.Encode(h.alerts); err != nil {
		log.Printf("error encoding messages: %v", err)
	}
}

// post receives the webhook from alertmanager and updates the AlertState
func (h *Hook) post(w http.ResponseWriter, req *http.Request) {
	log.Printf("received alert hook %v", req)

	msgBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("error reading message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}

	var msg *alert.Message
	err = json.Unmarshal(msgBytes, msg)
	if err != nil {
		log.Printf("error unmarshalling message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}

	h.Lock()
	defer h.Unlock()
	state := &AlertState{Msg: msg, Received: time.Now()}
	h.alerts[msg.Receiver] = state
}
