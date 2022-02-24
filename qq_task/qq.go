package qq_task

import (
	"fmt"
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
	log.Println(robotgo.GetMainDPI(),robotgo.GetHWND(),robotgo.GetPID())
	//mdpi :=robotgo.GetMainDPI()
	mpid := robotgo.GetPID()

	ids, err := robotgo.FindIds("QQPCTray")
	if err != nil {
		println(err)
		return
	}
	if len(ids) == 0 {
		println("没有检测到QQ电脑管家进程")
		return
	}
	log.Println(ids)
	robotgo.MoveClick(util.Width,util.Height)
	log.Println(robotgo.GetMainDPI(),robotgo.GetHWND(),robotgo.GetHandle())


	bit := robotgo.CaptureScreen(0, 0, util.Width/2, util.Height-util.WINDOWS_TASKBAR)
	defer robotgo.FreeBitmap(bit)
	robotgo.SaveBitmap(bit, util.PC_DNGJ_FILE_NAME)

	// Load full image
	img,_ := util.LoadImage(util.PC_DNGJ_FILE_NAME)

	// Create a lookup for that image
	l := lookup.NewLookup(img)

	// Load a template to search inside the image
	template,_ := util.LoadImage("images/dngj_font.png")

	// Find all occurrences of the template in the image
	pp, _ := l.FindAll(template, 0.9)

	// Print the results
	if len(pp) > 0 {
		//fmt.Printf("Found %d matches:\n", len(pp))
		for _, p := range pp {
			fmt.Printf("- (%d, %d) with %f accuracy\n", p.X, p.Y, p.G)
			//启动页面
			robotgo.MoveClick(p.X, p.Y,"left",true)
			//dngjPid := robotgo.GetPID()
			//dngj := robotgo.GetHWND()
			//切换移动鼠标
			robotgo.ActivePID(mpid)
			robotgo.MoveSmooth(util.Width/2-380, util.Height/2-230)


			//log.Println(robotgo.ActivePID(dngjPid))
			//robotgo.SetHandle(int(dngj))
			//log.Println(robotgo.SetActiveWindow(dngj))
			//log.Println(robotgo.ActiveName("电脑管家"))
			//log.Println("移动后:",robotgo.GetMainDPI(),robotgo.GetHWND(),robotgo.GetPID())
			robotgo.Click("left")
			log.Println("点击签到后:",robotgo.GetMainDPI(),robotgo.GetHWND(),robotgo.GetPID())
			robotgo.Click("left")

		}
	} else {
		println("No matches found")
	}
}
