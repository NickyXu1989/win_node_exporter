package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/mem"
)

type MemoryStatsCollector struct {
	memoryStatus *prometheus.Desc
}

func NewMemoryStatsCollector() *MemoryStatsCollector {
	return &MemoryStatsCollector{
		memoryStatus: prometheus.NewDesc("win_node_memory_status","the status of memory",[]string{"mode"},nil),
	}
}

func (collector *MemoryStatsCollector)Describe (ch chan <- *prometheus.Desc) {
	ch <- collector.memoryStatus
}

func (collector *MemoryStatsCollector)Collect (ch chan <- prometheus.Metric ) {
	//get memory status
	v, _ := mem.VirtualMemory()
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Total),"total")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Available),"available")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Used),"used")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.UsedPercent),"usedPercent")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Free),"free")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Active),"active")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Inactive),"inactive")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Wired),"wired")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Laundry),"laundry")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Buffers),"buffers")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Cached),"cached")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Writeback),"writeback")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Dirty),"dirty")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.WritebackTmp),"writebacktmp")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Shared),"shared")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Slab),"slab")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.SReclaimable),"sreclaimable")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.PageTables),"pagetables")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.CommitLimit),"commitlimit")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.CommittedAS),"committedas")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.HighTotal),"hightotal")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.HighFree),"highfree")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.LowTotal),"lowtotal")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.LowFree),"lowfree")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.SwapTotal),"swaptotal")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.SwapFree),"swapfree")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.Mapped),"mapped")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.VMallocTotal),"vmalloctotal")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.VMallocUsed),"vmallocused")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.VMallocChunk),"vmallocchunk")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.HugePagesTotal),"hugepagestotal")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.HugePagesFree),"hugepagesfree")
	ch <- prometheus.MustNewConstMetric(collector.memoryStatus,prometheus.GaugeValue,float64(v.HugePageSize),"hugepagesize")

}


