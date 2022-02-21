package qq_task

import (
	"github.com/deluan/lookup"
	"github.com/fengweiqiang/NBscripts/util"
	"github.com/go-vgo/robotgo"
	"log"
)

//QQ打卡
func PcQqDailyAttendance() {

	robotgo.Move(util.Width, 0)

	robotgo.MoveRelative(-150, 5)
	// 获取当前鼠标所在的位置
	x, y := robotgo.GetMousePos()

	robotgo.Click()

	robotgo.Move(util.Width/2, util.Height/2)
	// 获取当前鼠标所在的位置
	x, y = robotgo.GetMousePos()
	//println(`x：`, x, ` y：`, y)

	robotgo.MoveRelative(0, 90)
	robotgo.Click()
	println("签到成功,准备关闭签到页面")

	robotgo.Move(x+150, y-300)
	robotgo.Click()
	println("关闭签到页面")
}

//QQ腾讯管家打卡  获取不到窗口！！！
func PcQqTXGJDailyAttendance() {
	ids, err := robotgo.FindIds("QQPCTray")
	if err != nil {
		println(err)
		return
	}
	if len(ids) == 0 {
		println("没有检测到QQ电脑管家")
		return
	}
	//mainPId := ids[0]
	bit := robotgo.CaptureScreen(0, util.Height-util.WINDOWS_TASKBAR, util.Width, util.WINDOWS_TASKBAR)
	defer robotgo.FreeBitmap(bit)
	robotgo.SaveBitmap(bit, util.PC_DNGJ_FILE_NAME)

	image, err := util.LoadImage(util.PC_DNGJ_FILE_NAME)
	if err != nil {
		return
	}
	lookImage := lookup.NewLookup(image)
	dngj_ioc, err := util.LoadImage("images/dngj_ico.png")
	if err != nil {
		println(err)
		return
	}

	pp, err := lookImage.FindAll(dngj_ioc, 0.9)
	if err != nil {
		println(err)
		return
	}
	// Print the results
	if len(pp) > 0 {
		//log.Printf("Found %d matches:\n", len(pp))
		for _, p := range pp {
			//log.Printf("- (%d, %d) with %f accuracy\n", p.X, p.Y, p.G)
			robotgo.Move(p.X+15, util.Height-p.Y)

			robotgo.Click("left")

			//ids, err := robotgo.FindIds("QQPCTray")
			//if err != nil {
			//	println(err)
			//	return
			//}
			//for _, v := range ids {
			//	if mainPId == v {
			//		continue
			//	}
			//	err := robotgo.ActivePID(v)
			//	if err != nil {
			//		println(err)
			//		continue
			//	}

				x, y := robotgo.GetMousePos()
				log.Println(x,y)

				robotgo.MoveSmooth(util.Width/2-380, util.Height/2-230)

				x, y = robotgo.GetMousePos()
				log.Println(x,y)

				robotgo.Click("left", true)

				robotgo.Click()
				println("点击头像完成")
			//
			//}
		}
	} else {
		println("No matches found")
	}
}
