package common

import (
	"fmt"
	"os"
	"os/exec"
)

type ClearCmd struct{}

func (c *ClearCmd) ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("控制台清空失败:", err)
		return
	}
}
