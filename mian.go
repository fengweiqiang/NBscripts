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
		case 2:
			println("执行PC端管家签到")
			qq_task.PcQqTXGJDailyAttendance()
	}


	return
	//robotgo.MouseSleep = 2000
	//robotgo.MoveClick(width, 6)
	//robotgo.MouseSleep = 2000
	//robotgo.MoveClick(1277, 911)
	//println("QQ签到成功")
	//robotgo.MoveClick(1431, 436)
	//println("QQ签到页面关闭")
	//// 获取当前鼠标所在的位置
	//x, y = robotgo.GetMousePos()
	//println(`x：`, x, ` y：`, y)
	//return
	//
	//robotgo.Click("wheelRight")
	//
	//robotgo.Click("left", true)
	//
	//robotgo.Toggle("left")
	//robotgo.Toggle("left", "up")
}

//func opencv() {
//	name := "test.png"
//	name1 := "test_001.png"
//	robotgo.SaveCapture(name1, 10, 10, 30, 30)
//	robotgo.SaveCapture(name)
//
//	fmt.Print("gcv find image: ")
//	fmt.Println(gcv.FindImgFile(name1, name))
//	fmt.Println(gcv.FindAllImgFile(name1, name))
//
//	bit := bitmap.Open(name1)
//	defer robotgo.FreeBitmap(bit)
//	fmt.Print("find bitmap: ")
//	fmt.Println(bitmap.Find(bit))
//
//	// bit0 := robotgo.CaptureScreen()
//	// img := robotgo.ToImage(bit0)
//	// bit1 := robotgo.CaptureScreen(10, 10, 30, 30)
//	// img1 := robotgo.ToImage(bit1)
//	// defer robotgo.FreeBitmapArr(bit0, bit1)
//	img := robotgo.CaptureImg()
//	img1 := robotgo.CaptureImg(10, 10, 30, 30)
//
//	fmt.Print("gcv find image: ")
//	fmt.Println(gcv.FindImg(img1, img))
//	fmt.Println()
//
//	res := gcv.FindAllImg(img1, img)
//	fmt.Println(res[0].TopLeft.Y, res[0].Rects.TopLeft.X, res)
//	x, y := res[0].TopLeft.X, res[0].TopLeft.Y
//	robotgo.Move(x, y-rand.Intn(5))
//	robotgo.MilliSleep(100)
//	robotgo.Click()
//
//	res = gcv.FindAll(img1, img) // use find template and sift
//	fmt.Println("find all: ", res)
//	res1 := gcv.Find(img1, img)
//	fmt.Println("find: ", res1)
//
//	img2, _, _ := robotgo.DecodeImg("test_001.png")
//	x, y = gcv.FindX(img2, img)
//	fmt.Println(x, y)
//}
