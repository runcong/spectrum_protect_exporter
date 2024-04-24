package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var db_total_fs_size_mb = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_db_total_fs_size_mb",
	Help: "Spectrum protect db total fs size in MB",
}, []string{"location"})

var db_used_fs_size_mb = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_db_used_fs_size_mb",
	Help: "Spectrum protect db used fs size in MB",
}, []string{"location"})

var db_free_space_mb = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_db_free_space_mb",
	Help: "Spectrum protect db free space in MB",
}, []string{"location"})

func db_status() {
	output := dsmadmc_query("select location, total_fs_size_mb, used_fs_size_mb, free_space_mb from dbspace")
	// Split the output into lines
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line != "" {
			location := strings.Fields(line)[0]

			totalSize, err := strconv.ParseFloat(strings.Fields(line)[1], 64)
			if err != nil {
				fmt.Println("Error parsing otalSize:", err)
				continue
			}
			db_total_fs_size_mb.WithLabelValues(location).Set(totalSize)

			usedSize, err := strconv.ParseFloat(strings.Fields(line)[2], 64)
			if err != nil {
				fmt.Println("Error parsing usedSize:", err)
				continue
			}
			db_used_fs_size_mb.WithLabelValues(location).Set(usedSize)

			freeSize, err := strconv.ParseFloat(strings.Fields(line)[3], 64)
			if err != nil {
				fmt.Println("Error parsing freeSize:", err)
				continue
			}
			db_free_space_mb.WithLabelValues(location).Set(freeSize)

		}
	}

}
