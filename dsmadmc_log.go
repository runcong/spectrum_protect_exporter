package main

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_protect_active_log_space = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_active_log_space",
	Help: "Spectrum Protect active log space in MB",
}, []string{"status"})

var spectrum_protect_archive_log_fs = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_archive_log_fs",
	Help: "Spectrum Protect archive log fs in MB",
}, []string{"status"})

func get_log_usage(status string) float64 {
	// query := "SELECT pct_utilized FROM stgpools WHERE stgpool_name='" + status + "'"
	// query := "select" + status + "from log"
	query := "select " + status + " from log"
	// fmt.Printf("query: %s\n", query)
	output := dsmadmc_query(query)
	// fmt.Printf("output: %s\n", output)
	if output == "" {
		// fmt.Println("Error getting pct_utilized for pool", pool_name)
		log.Fatal("Error get_active_log_space", status)
		return 0
	} else {
		value, err := strconv.ParseFloat(output, 64)
		if err != nil {
			log.Fatal(err)
			// fmt.Println("Error getting pct_utilized for pool", pool_name, err)
			return 0
		} else {
			return value
		}
	}
}

func active_log_space() {
	status_s := []string{"total_space_mb", "used_space_mb", "free_space_mb"}
	for _, status := range status_s {
		value := get_log_usage(status)
		spectrum_protect_active_log_space.WithLabelValues(status).Set(value)
	}
}

func archive_log_fs() {
	status_s := []string{"archlog_tol_fs_mb", "archlog_used_fs_mb", "archlog_free_fs_mb"}
	for _, status := range status_s {
		value := get_log_usage(status)
		spectrum_protect_archive_log_fs.WithLabelValues(status).Set(value)
	}
}
