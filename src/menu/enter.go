package menu

import (
	"GoLearn/src/common"
	"GoLearn/src/shutdown"
	"fmt"
)

type Strategy interface {
	Execute()
}

type Menu struct{}

func (m *Menu) Display() int {
	clearCMD()
	var choice int
	fmt.Println("请选择操作:")
	fmt.Println("1. 定时关机")
	fmt.Println("2. 其他操作（待添加）")
	fmt.Print("输入选项: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0
	}
	return choice
}

func (m *Menu) SelectStrategy(choice int) Strategy {
	clearCMD()
	switch choice {
	case 1:
		return &shutdown.BeginShutdownStrategy{}
	default:
		return nil
	}
}

// 清空控制台
func clearCMD() {
	clearCmd := common.ClearCmd{}
	clearCmd.ClearConsole()
}
