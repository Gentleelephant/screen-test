package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	// 模拟鼠标点击
	simulateLeftMouseClick()

}

func simulateLeftMouseClick() {
	robotgo.Click("left", false)

	// 等待一段时间，可以省略
	time.Sleep(1 * time.Second)

	// 模拟释放左键
	robotgo.Click("left", true)

	fmt.Println("toggle up", robotgo.Toggle("up"))
}
