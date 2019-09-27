package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"log"
	"net/http"
	"win_node_exporter/util"


	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//"time"

	//"github.com/shirou/gopsutil/disk"
)

func init()  {
	//prometheus.MustRegister(util.CpuCount)
	//prometheus.MustRegister(util.CpuTimes)
	test()

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



func test() {
	fmt.Println("memory")
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Println("CPU")
	c, _ := cpu.Counts(true)
	t,_ := cpu.Times(true)
	fmt.Println(c)
	fmt.Println(t)


	//drivers := GetLogicalDrives()
	//fmt.Println(drivers)
	//d, _ := disk.IOCounters()
	drivers, _ := disk.Partitions(true)
	fmt.Println("partitions:")
	fmt.Println(drivers)
	fmt.Println()

	fmt.Println("IOCounters")
	for _,driver:= range drivers {
		ds, _ := disk.IOCounters(driver.Device)
		fmt.Println(ds)

		for k,v := range ds {
			fmt.Println(k)
			fmt.Println(v)
		}

		fmt.Println("disk usage")
		du,_ := disk.Usage(driver.Device)
		fmt.Println(du)
	}



	fmt.Println("net")
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces)

	for _,i := range interfaces {
		fmt.Println(i.Name)
		fmt.Println(net.IOCountersByFile(true,i.Name))
	}

	//iocounter, _ := net.IOCounters(true)
	//fmt.Println(iocounter)

}
//
//
//func GetLogicalDrives() []string {
//	kernel32 := syscall.MustLoadDLL("kernel32.dll")
//	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
//	n, _, _ := GetLogicalDrives.Call()
//	s := FormatInt(int64(n), 2)
//	var drivesAll = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
//	temp := drivesAll[0:len(s)]
//	var d []string
//	for i, v := range s {
//		if v == 49 {
//			l := len(s) - i - 1
//			d = append(d, temp[l])
//		}
//	}
//	var drives []string
//	for i, v := range d {
//		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
//	}
//	return drives
//}