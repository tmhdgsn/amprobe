package main

import (
	"flag"
	"log"

	"github.com/tmhdgsn/amprobe/hook"
)

func main() {
	addr := flag.String("addr", "8080", "listen addr for webhook")
	flag.Parse()

	hook := hook.New(*addr)
	if err := hook.ListenAndServe(); err != nil {
		log.Fatalf("err: %s", err)
	}
}
