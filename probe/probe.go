package probe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prometheus/common/model"
)

type (
	Probe struct {
		BaseURL *url.URL
		Client  *http.Client
	}

	Alerts []model.Alert
)

func New(httpClient *http.Client) *Probe {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Probe{
		BaseURL: &url.URL{Scheme: "http", Host: "localhost:9093"},
		Client:  httpClient,
	}
}

func (p *Probe) SendAlerts(alerts Alerts) error {
	rel := &url.URL{Path: "/api/v1/alerts"}
	u := p.BaseURL.ResolveReference(rel)
	data, err := json.Marshal(alerts)
	if err != nil {
		return err
	}

	payload := bytes.NewReader(data)
	fmt.Printf("alert: %s\n", data)
	req, err := http.NewRequest("POST", u.String(), payload)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("RESP: %+v\n", resp)

	return nil
}
