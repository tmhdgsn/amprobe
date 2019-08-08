package probe

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/common/model"
)

type (
	Probe struct {
		BaseURL *url.URL
		Client  *http.Client
		ticker  *time.Ticker
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
		ticker:  time.NewTicker(30 * time.Second),
	}
}

func (p *Probe) SendAlerts(alerts Alerts) error {
	rel := &url.URL{Path: "/api/v1/alerts"}
	u := p.BaseURL.ResolveReference(rel)
	data, err := json.Marshal(alerts)
	if err != nil {
		return err
	}

	for {
		select {
		case <-p.ticker.C:
			payload := bytes.NewReader(data)
			req, err := http.NewRequest("POST", u.String(), payload)
			if err != nil {
				return err
			}

			resp, err := p.Client.Do(req)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				bytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				body := string(bytes)
				log.Print(body)
			}

		}
	}

	return nil
}
