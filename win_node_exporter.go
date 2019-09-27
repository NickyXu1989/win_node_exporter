package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"win_node_exporter/util"
)

func init()  {

}

func registerCpuMetrics() {
	cpuCollector := util.NewCpuStatsCollector()
	prometheus.MustRegister(cpuCollector)
}

func registerDiskMetrics() {
	diskCollector := util.NewDiskStatsCollector()
	prometheus.MustRegister(diskCollector)
}

func registerMemoryMetrics() {
	memoryCollector := util.NewMemoryStatsCollector()
	prometheus.MustRegister(memoryCollector)
}

func registerNetMetrics() {
	netCollector := util.NewNetStatsCollector()
	prometheus.MustRegister(netCollector)
}

func main()  {

	registerCpuMetrics()
	registerDiskMetrics()
	registerMemoryMetrics()
	registerNetMetrics()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8099", nil))
}


