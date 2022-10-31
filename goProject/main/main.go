package main

import (
	"fmt"
	"go-demo/goProject/demo02"
	"go-demo/goProject/function"
	"time"
)

func main() {
	fmt.Println(function.TwoNumAdd(12, 13))
	fmt.Println(function.AllNumConvertBIN(3))
	go demo02.PrintfHelloWorld()
	time.Sleep(time.Second)
	function.LeapYearJudge(2022)
	function.SwitchJudge(86)
	function.SingularJudge(6)
	function.Test(4)
	fmt.Printf("%T\n", function.Test2(0))
	function.GetCpuInformation() //获取CPU总体信息
	fmt.Printf("CPU使用率: %0.1f%%\n内存信息: %v%%\n磁盘使用容量: %0.1f%%\n", function.GetCpuPercent(), function.GetMemPercent(), function.GetDiskPercent())
	go function.GetCpuUtilization()
	time.Sleep(time.Hour)
}
