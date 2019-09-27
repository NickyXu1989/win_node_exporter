package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"strings"
)

type CpuStatsCollector struct {
	cpuCounts *prometheus.Desc
	cpuTimes *prometheus.Desc
}

func NewCpuStatsCollector() *CpuStatsCollector {
	return &CpuStatsCollector{
		cpuCounts: prometheus.NewDesc("win_node_cpu_counts","total nums of cpus",nil,nil),
		cpuTimes:  prometheus.NewDesc("win_node_cpu_times","cpu times of each cpu", []string{"cpu_id","mode"}, nil),
	}
}

func (collector *CpuStatsCollector)Describe (ch chan <- *prometheus.Desc) {
	ch <- collector.cpuCounts
	ch <- collector.cpuTimes
}

func (collector *CpuStatsCollector)Collect (ch chan <- prometheus.Metric ) {

	//get cpu counts
	cpuCount := getCpuCount()
	ch <- prometheus.MustNewConstMetric(collector.cpuCounts, prometheus.GaugeValue,cpuCount,)

	// get cpu times
	cpuTimes, _ := cpu.Times(true)
	for _, cpuTime := range cpuTimes {
		//count cpu id
		cpuIdStr := cpuTime.CPU
		cpuIdArr := strings.Split(cpuIdStr, ",")
		cpuSocketId,_ := strconv.Atoi(cpuIdArr[0])
		cpuCoreId, _:= strconv.Atoi(cpuIdArr[1])
		cpuId := (cpuSocketId + 1) * cpuCoreId
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.User,strconv.Itoa(cpuId),"user")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.System,strconv.Itoa(cpuId),"system")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Idle,strconv.Itoa(cpuId),"idle")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Nice,strconv.Itoa(cpuId),"nice")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Iowait,strconv.Itoa(cpuId),"iowait")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Irq,strconv.Itoa(cpuId),"irq")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Softirq,strconv.Itoa(cpuId),"softirq")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Steal,strconv.Itoa(cpuId),"steal")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Guest,strconv.Itoa(cpuId),"guest")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.GuestNice,strconv.Itoa(cpuId),"guestNice")
		ch <- prometheus.MustNewConstMetric(collector.cpuTimes, prometheus.CounterValue, cpuTime.Stolen,strconv.Itoa(cpuId),"stolen")
	}

}



func getCpuCount() float64{
	c, _ := cpu.Counts(true)
	return (float64(c))
}