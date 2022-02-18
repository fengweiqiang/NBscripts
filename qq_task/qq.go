package qq_task

import (
	"github.com/fengweiqiang/NBscripts/util"
	"github.com/go-vgo/robotgo"
)

//QQ打卡
func PcQqDailyAttendance()  {

	robotgo.MouseSleep = 1000
	robotgo.Move(util.Width, 0)
	robotgo.MouseSleep = 1000

	robotgo.MoveRelative(-150, 5)
	robotgo.MouseSleep = 1000
	// 获取当前鼠标所在的位置
	x, y := robotgo.GetMousePos()
	robotgo.MouseSleep = 1000

	robotgo.Click()
	robotgo.MouseSleep = 1000
	robotgo.Move(util.Width/2, util.Height/2)
	robotgo.MouseSleep = 1000
	// 获取当前鼠标所在的位置
	x, y = robotgo.GetMousePos()
	robotgo.MouseSleep = 1000
	//println(`x：`, x, ` y：`, y)

	robotgo.MoveRelative(0, 90)
	robotgo.MouseSleep = 2000
	robotgo.Click()
	println("签到成功,准备关闭签到页面")

	robotgo.Move(x+150, y-300)
	robotgo.MouseSleep = 500
	robotgo.Click()
	println("关闭签到页面")
}
//QQ腾讯管家打卡
func PcQqTXGJDailyAttendance()  {

}
