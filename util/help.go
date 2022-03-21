package util

import (
	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	"github.com/shopspring/decimal"
	"image"
	"log"
	"os"
	"syscall"
	"unsafe"
)

var Width, Height int

const (
	//任务栏图标宽度
	WINDOWS_ICON_WIDTH = 30
	//任务栏高度
	WINDOWS_TASKBAR   = 50
	PC_DNGJ_FILE_NAME = "images/dngj.png"
	PC_ZN_FILE_NAME   = "images/zm.png"
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

var (
	moduser32                 = syscall.NewLazyDLL("user32.dll")
	procGetMonitorInfo        = moduser32.NewProc("GetMonitorInfoW")
	procEnumDisplaySettingsEx = moduser32.NewProc("EnumDisplaySettingsExW")
)

const (
	ENUM_CURRENT_SETTINGS  = 0xFFFFFFFF
	ENUM_REGISTRY_SETTINGS = 0xFFFFFFFE
)

//get system dpi
func GetDpiForSystem() (horzScalestr, vertScalestr float64) {
	// 获取窗口当前显示的监视器
	hWnd := win.MonitorFromWindow(win.GetConsoleWindow(), win.MONITOR_DEFAULTTONEAREST)
	// 获取监视器逻辑宽度与高度
	var mi win.MONITORINFOEX
	mi.CbSize = uint32(unsafe.Sizeof(mi))
	getMonitorInfoEx(hWnd, &mi)
	cxLogical := (mi.RcMonitor.Right - mi.RcMonitor.Left)
	cyLogical := (mi.RcMonitor.Bottom - mi.RcMonitor.Top)
	// 获取监视器物理宽度与高度
	var dm win.DEVMODE
	dm.DmSize = uint16(unsafe.Sizeof(dm))
	dm.DmDriverExtra = 0

	enumDisplaySettingsEx(&mi.SzDevice[0], ENUM_CURRENT_SETTINGS, &dm, 0)
	cxPhysical := dm.DmPelsWidth
	cyPhysical := dm.DmPelsHeight

	//缩放比例计算
	horzScale, _ := decimal.NewFromInt32(int32(cxPhysical)).Div(decimal.NewFromInt32(int32(cxLogical))).Round(2).Float64()
	vertScale, _ := decimal.NewFromInt32(int32(cyPhysical)).Div(decimal.NewFromInt32(int32(cyLogical))).Round(2).Float64()
	return horzScale, vertScale
}

func getMonitorInfoEx(hMonitor win.HMONITOR, lmpi *win.MONITORINFOEX) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}
func enumDisplaySettingsEx(szDeviceName *uint16, iModeNum uint32, devMode *win.DEVMODE, dwFlags uint32) bool {
	ret, _, _ := procEnumDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(dwFlags),
	)
	return ret != 0
}
