package main

import (
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/model"

	"github.com/tmhdgsn/amprobe/hook"
	"github.com/tmhdgsn/amprobe/probe"
)

func main() {

	amprobe := probe.New(nil)
	amhook := hook.New("1337")

	alert := model.Alert{
		StartsAt:    time.Now(),
		EndsAt:      time.Now().Add(time.Duration(30 * time.Second)),
		Labels:      model.LabelSet{"label1": "test1"},
		Annotations: model.LabelSet{"annotation1": "some text"},
	}

	alerts := []model.Alert{alert}

	if err := amprobe.SendAlerts(alerts); err != nil {
		log.Fatalf("err: %s", err)
	}

	if err := amhook.ListenAndServe(promhttp.Handler()); err != nil {
		log.Fatalf("err: %s", err)
	}

	fmt.Printf("alert: %+v\n", alert)
}
