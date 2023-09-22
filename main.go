package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

func main() {
	captureScreen(0)
}

func captureScreen(displayIndex int) {

	// 获取屏幕的尺寸
	bounds := screenshot.GetDisplayBounds(displayIndex)

	// 创建一个图像对象，用于存储屏幕内容
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		fmt.Println("无法捕获屏幕内容:", err)
		return
	}

	// 创建一个文件来保存图像内容
	file, err := os.Create(fmt.Sprintf("screenshot%d.png", displayIndex))
	if err != nil {
		fmt.Println("无法创建图像文件:", err)
		return
	}
	defer file.Close()
	//
	// 将图像编码为PNG格式并保存到文件中
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("无法保存图像:", err)
	} else {
		fmt.Println(fmt.Sprintf("保存成功:screenshot%d.png", displayIndex))
	}

	judgeCenterColor(img)

}

func judgeCenterColor(img image.Image) bool {
	// 获取图像的宽度和高度
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// 计算图像中心坐标
	centerX := width / 2
	centerY := height / 2

	// 获取图像中心像素颜色
	centerColor := img.At(centerX, centerY)

	//// 判断颜色是否为红色
	if isRed(centerColor) {
		fmt.Println("颜色符合")
	} else {
		fmt.Println("颜色不符合")
	}

	return false
}

// 判断颜色是否为红色
func isRed(c color.Color) bool {
	r, _, _, _ := c.RGBA()
	// 判断红色分量是否大于阈值

	// 判断红色分量是否大于阈值
	return r > 0x8000
}

// 将颜色值字符串解析为 RGBA 颜色
func hexToRGBA(hex string) (color.RGBA, error) {

	// 定义变量用于存储RGB分量
	var r, g, b uint8
	// 去掉HEX颜色字符串中的'#'字符
	hex = strings.TrimPrefix(hex, "#")
	// 确保HEX字符串至少有6个字符
	if len(hex) < 6 {
		return color.RGBA{}, fmt.Errorf("无效的HEX颜色字符串: %s", hex)
	}
	// 解析RGB分量
	_, err := fmt.Sscanf(hex[:6], "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return color.RGBA{}, err
	}
	// 创建RGBA颜色，Alpha通道为255表示完全不透明
	rgbaColor := color.RGBA{R: r, G: g, B: b, A: 255}
	return rgbaColor, nil
}
