package util

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"
)

type NetStatsCollector struct {
	netStatus *prometheus.Desc
}



func NewNetStatsCollector() *NetStatsCollector {
	return &NetStatsCollector{
		netStatus: prometheus.NewDesc("win_node_net_status","status of interfaces",[]string{"name","mode"},nil),
	}
}

func (collector *NetStatsCollector)Describe (ch chan <- *prometheus.Desc) {
	ch <- collector.netStatus
}

func (collector *NetStatsCollector)Collect (ch chan <- prometheus.Metric ) {

	tmpSet := mapset.NewSet()

	//get all interfaces
	interfaces, _ := net.Interfaces()
	for _,i := range interfaces {
		fmt.Println(net.IOCountersByFile(true,i.Name))
		interfaceStatus,_ := net.IOCountersByFile(true, i.Name)
		for _,iStauts := range interfaceStatus {
			// in case of collecting duplicated interfaces
			if tmpSet.Contains(iStauts.Name) {
				continue
			}
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.BytesSent),iStauts.Name,"bytesSent")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.BytesRecv),iStauts.Name,"bytesRecv")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.PacketsSent),iStauts.Name,"packetsSent")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.PacketsRecv),iStauts.Name,"packetsRecv")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Errin),iStauts.Name,"errin")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Errout),iStauts.Name,"errout")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Dropin),iStauts.Name,"dropin")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Dropout),iStauts.Name,"dropout")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Fifoin),iStauts.Name,"fifoin")
			ch <- prometheus.MustNewConstMetric(collector.netStatus, prometheus.GaugeValue,float64(iStauts.Fifoout),iStauts.Name,"fifoout")
			tmpSet.Add(iStauts.Name)
		}
	}

}



func getMetCount() float64{
	c, _ := cpu.Counts(true)
	return (float64(c))
}