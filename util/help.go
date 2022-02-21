package util

import (
	"github.com/go-vgo/robotgo"
	"image"
	"log"
	"os"
)

var Width, Height int

const (
	//任务栏图标宽度
	WINDOWS_ICON_WIDTH = 30
	//任务栏高度
	WINDOWS_TASKBAR   = 50
	PC_DNGJ_FILE_NAME = "images/dngj.png"
)

func init() {
	robotgo.MouseSleep = 1000
	robotgo.KeySleep = 1000
}

func GetResolvingPower() {
	// 获取屏幕分辨率
	width, height := robotgo.GetScreenSize()
	println("获取屏幕分辨率", `x宽：`, width, ` y高：`, height)
	if width == 0 || height == 0 {
		panic("获取屏幕分辨率异常")
	}
	if width <= 800 && height <= 600 {
		println("当前屏幕分辨率不支持,请使用其他分辨率")
	}
	Width = width
	Height = height

}

// Helper function to load an image from the filesystem
func LoadImage(imgPath string) (image.Image, error) {
	imageFile, err := os.Open(imgPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img, nil
}
