package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func generateLogs() {
	sum := 1
	for sum < 1000 {
		log.Println("Log number", sum)
		sum += 1
		time.Sleep(15 * time.Second)
	}

}

func main() {
	go generateLogs()

	reg := prometheus.NewRegistry()

	reg.MustRegister()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
		log.Println("Health log")
	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})

	http.ListenAndServe(":9999", nil)
}
