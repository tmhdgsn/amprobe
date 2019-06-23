package main

import (
	"fmt"
	"log"
	"time"

	"github.com/prometheus/common/model"

	"github.com/tmhdgsn/amprobe/probe"
)

func main() {

	amprobe := probe.New(nil)

	alert := model.Alert{
		StartsAt:    time.Now(),
		EndsAt:      time.Now().Add(time.Duration(5 * time.Minute)),
		Labels:      model.LabelSet{"label1": "test1"},
		Annotations: model.LabelSet{"annotation1": "some text"},
	}

	alerts := []model.Alert{alert}

	if err := amprobe.SendAlerts(alerts); err != nil {
		log.Fatalf("err: %s", err)
	}

	if err := amprobe.Hook.ListenAndServe(); err != nil {
		log.Fatalf("err: %s", err)
	}

	fmt.Printf("alert: %+v\n", alert)
}
