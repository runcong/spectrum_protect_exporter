package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_protect_admin_schedule = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_admin_schedule",
	Help: "Spectrum protect admin schedule task status",
}, []string{"schedule_start", "schedule_name", "status"})

func admin_schedule() {
	querry := "select scheduled_start,actual_start,schedule_name,status from events where scheduled_start>current_timestamp - 24 hours and domain_name is NULL order by scheduled_start"
	output := dsmadmc_query(querry)
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			length := len(strings.Fields(line))
			schedule_start := strings.Fields(line)[0] + " " + strings.Fields(line)[1]
			schedule_name := strings.Fields(line)[length-2]
			status := strings.Fields(line)[length-1]
			spectrum_protect_admin_schedule.WithLabelValues(schedule_start, schedule_name, status).Set(0)
		}
	}

}
