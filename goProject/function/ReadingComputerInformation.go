package function

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// GetCpuPercent
// @description 获取CPU占用率
// @return float64
func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

// GetMemPercent
// @description 获取内存信息
// @return float64
func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

// GetDiskPercent
// @description 获取磁盘使用容量百分比
// @return float64
func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

// GetNetworkInformation
// @description 获取当前网络连接信息
func GetNetworkInformation() {
	info, _ := net.Connections("tcp") //可填入tcp、udp、tcp4、udp4等等
	fmt.Println(info)
}

// GetCpuUtilization
// @description 获取CPU占用率,1秒刷新一次
func GetCpuUtilization() {
	for {
		info, _ := cpu.Percent(time.Duration(time.Second*3), false)
		fmt.Printf("CPU占用率:%0.1f\n", info)
	}
}

// GetCpuInformation
// @description 获取CPU总体信息
func GetCpuInformation() {
	info, _ := cpu.Info()
	fmt.Println(info)
}
