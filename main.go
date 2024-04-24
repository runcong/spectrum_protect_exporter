package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var cmdName string = "dsmadmc"
var userID string = "dstadmin"

// suply your password
var password string = "YOUR_PASSWORD"
var dataOnly string = "yes"

var listenAddress = flag.String(
	"listen-address",
	":9109",
	"The address to listen on for HTTP requests.")

var reg = prometheus.NewRegistry()

func dsmadmc_query(query string) string {
	cmd := exec.Command(cmdName, "-id="+userID, "-password="+password, "-dataonly="+dataOnly, query)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error running dsmadmc: %s\n", err)
		return "0"
	}
	if len(output) == 0 {
		log.Fatal("No output from dsmadmc")
		fmt.Println("No output from dsmadmc")
		return "0"
	}
	return strings.TrimSpace(string(output))
}

func registerMetrics() {
	reg.MustRegister(spectrum_protect_tapes)
	reg.MustRegister(spectrum_protect_pct_utilized)
	reg.MustRegister(spectrum_protect_active_log_space)
	reg.MustRegister(spectrum_protect_archive_log_fs)
	reg.MustRegister(db_total_fs_size_mb)
	reg.MustRegister(db_used_fs_size_mb)
	reg.MustRegister(db_free_space_mb)
	reg.MustRegister(spectrum_protect_admin_schedule)

}

func main() {
	registerMetrics()
	flag.Parse()
	log.Printf("Starting Server: %s", *listenAddress)
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go tapes_number()
		go pct_utilized()
		go active_log_space()
		go archive_log_fs()
		go db_status()
		go admin_schedule()

		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	}))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
