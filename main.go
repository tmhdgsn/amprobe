package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/common/model"

	"github.com/tmhdgsn/amprobe/config"
	"github.com/tmhdgsn/amprobe/hook"
	"github.com/tmhdgsn/amprobe/probe"
)

var (
	flgAlertRule = flag.String("alertRule", "rules/alert-0.yaml", "alert rule file")
)

func main() {
	flag.Parse()

	ruleFiles := []string{*flgAlertRule}

	rules, errs := config.LoadRules(ruleFiles)
	if errs != nil {
		log.Fatalf("err: %s\n", errs[0])
	}

	fmt.Printf("loaded rules: %+v\n", rules)

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

	if err := amhook.ListenAndServe(); err != nil {
		log.Fatalf("err: %s", err)
	}

	fmt.Printf("alert: %+v\n", alert)
}
