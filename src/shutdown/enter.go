package shutdown

import (
	"fmt"
	"os/exec"
	"time"
)

type BeginShutdownStrategy struct{}

func (s *BeginShutdownStrategy) Execute() {
	var minutes int
	fmt.Println("-1. 返回菜单")
	fmt.Print("请输入关机时间(分钟): ")
	_, err := fmt.Scanln(&minutes)
	if err != nil {
		return
	}

	if minutes == -1 {
		return
	}

	duration := time.Duration(minutes) * time.Minute
	timer := time.NewTimer(duration)

	fmt.Printf("计算机将在 %v 后关闭\n", duration)

	ticker := time.NewTicker(1 * time.Second) // 定义每秒更新一次的 ticker
	defer ticker.Stop()

	start := time.Now()

	for {
		select {
		case <-timer.C:
			fmt.Println("时间到，正在关闭...")
			shutdown()
			return
		case <-ticker.C:
			elapsed := time.Since(start)
			remaining := duration - elapsed
			if remaining < 0 {
				remaining = 0
			}
			fmt.Printf("\r剩余时间: %v ", remaining.Round(time.Second))
		}
	}
}

func shutdown() {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println("关机失败:", err)
	} else {
		fmt.Println("计算机正在关闭...")
	}
}
