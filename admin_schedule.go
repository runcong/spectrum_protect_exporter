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
	// cmd := exec.Command("cat", "admin_schedule.txt")
	// output, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println("Error executing command:", err)
	// 	return
	// }
	querry := "select scheduled_start,actual_start,schedule_name,status from events where scheduled_start>current_timestamp - 24 hours and domain_name is NULL order by scheduled_start"
	output := dsmadmc_query(querry)
	// fmt.Printf("output: %s\n", output)
	// Split the output into lines
	lines := strings.Split(string(output), "\n")

	// Process each line
	for _, line := range lines {
		// Extract the tape status
		if line != "" {
			length := len(strings.Fields(line))
			// Extract the tape status
			schedule_start := strings.Fields(line)[0] + " " + strings.Fields(line)[1]
			// actual_start := strings.Fields(line)[3]
			schedule_name := strings.Fields(line)[length-2]
			status := strings.Fields(line)[length-1]
			spectrum_protect_admin_schedule.WithLabelValues(schedule_start, schedule_name, status).Set(0)
		}
	}

}
