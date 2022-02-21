package util

import "github.com/go-vgo/robotgo"

var Width,Height int

const ICON_WIDTH = 30

func GetResolvingPower() {
	// 获取屏幕分辨率
	width, height := robotgo.GetScreenSize()
	println("获取屏幕分辨率", `x宽：`, width, ` y高：`, height)
	if width == 0 || height == 0 {
		panic("获取屏幕分辨率异常")
	}
	Width = width
	Height = height

}
