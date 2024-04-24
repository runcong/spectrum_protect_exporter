package main

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_protect_tapes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_tapes",
	Help: "Spectrum Protect tapes number",
}, []string{"status"})

func tapes_number() {
	status := "scratch"
	query := "select count(*) from libvolumes where status='Scratch'"
	value, err := strconv.ParseFloat(dsmadmc_query(query), 32)
	if err != nil {
		log.Fatal(err)
	}
	spectrum_protect_tapes.WithLabelValues(status).Set(value)
}
