package main

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var spectrum_protect_pct_utilized = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "spectrum_protect_pct_utilized",
	Help: "Percentage of data in the pool",
}, []string{"pool_name"})

func get_pct_utilized(pool_name string) float64 {
	query := "SELECT pct_utilized FROM stgpools WHERE stgpool_name='" + pool_name + "'"
	output := dsmadmc_query(query)
	if output == "" {
		fmt.Println("getting empty value for pool:", pool_name)
		// log.Fatal("Error getting pct_utilized for pool", pool_name)
		// return 0
	} else {
		value, err := strconv.ParseFloat(output, 64)
		if err != nil {
			// log.Fatal(err)
			fmt.Println("Error getting pct_utilized for pool", pool_name, err)
			// return 0
		} else {
			return value
		}
	}
	return 0
}

func pct_utilized() {
	pool_names := []string{"TAPEPOOL", "FILEPOOL", "FUSION"}
	for _, pool_name := range pool_names {
		value := get_pct_utilized(pool_name)
		spectrum_protect_pct_utilized.WithLabelValues(pool_name).Set(value)
	}

}
