package main

import (
	"flag"
	"fmt"
	"github.com/fengweiqiang/NBscripts/qq_task"
	"github.com/fengweiqiang/NBscripts/util"
)

func main() {
	util.GetResolvingPower()
	var taskNum int
	println("简单脚本\n1：PC端QQ签到 2: PC端电脑管家签到")
	flag.Parse()
	println("请输入数字：")

	fmt.Scanln(&taskNum)
	switch taskNum {
	case 1:
		println("执行PC端QQ签到")
		qq_task.PcQqDailyAttendance()
	default:
		println("执行PC端管家签到")
		qq_task.PcQqTXGJDailyAttendance()
	}

	return
}
