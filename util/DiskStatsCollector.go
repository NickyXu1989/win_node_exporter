package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/disk"
)

type DiskStatsCollector struct {
	diskUsage *prometheus.Desc
	diskStatus *prometheus.Desc
}

func NewDiskStatsCollector() *DiskStatsCollector {
	return &DiskStatsCollector{
		diskUsage: prometheus.NewDesc("win_node_disk_usage","the usage of the disk",[]string{"path","mode"},nil),
		diskStatus:  prometheus.NewDesc("win_node_disk_status","the status of the disk", []string{"path","mode"}, nil),
	}
}

func (collector *DiskStatsCollector)Describe (ch chan <- *prometheus.Desc) {
	ch <- collector.diskUsage
	ch <- collector.diskStatus
}

func (collector *DiskStatsCollector)Collect (ch chan <- prometheus.Metric ) {

	//get all disks
	drivers, _ := disk.Partitions(true)

	for _,driver:= range drivers {
		//get disk io status
		ds, _ := disk.IOCounters(driver.Device)
		for k,v := range ds {
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.ReadCount),k,"readCount")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.MergedReadCount),k,"mergedReadCount")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.WriteCount),k,"writeCount")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.MergedWriteCount),k,"mergedWriteCount")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.ReadBytes),k,"readBytes")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.WriteBytes),k,"writeBytes")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.ReadTime),k,"readTime")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.WriteTime),k,"writeTime")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.IopsInProgress),k,"iopsInProgress")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.IoTime),k,"ioTime")
			ch <- prometheus.MustNewConstMetric(collector.diskStatus, prometheus.GaugeValue,float64(v.WeightedIO),k,"weightedIO")
		}

		//get disk usage
		du,_ := disk.Usage(driver.Device)
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.Total),driver.Device,"total")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.Free),driver.Device,"free")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.Used),driver.Device,"used")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.UsedPercent),driver.Device,"usedPercent")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.InodesTotal),driver.Device,"inodesTotal")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.InodesUsed),driver.Device,"inodesUsed")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.InodesFree),driver.Device,"inodesFree")
		ch <- prometheus.MustNewConstMetric(collector.diskUsage, prometheus.GaugeValue,float64(du.InodesUsedPercent),driver.Device,"inodesUsedPercent")

	}

}


