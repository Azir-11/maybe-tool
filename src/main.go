package main

import (
	"GoLearn/src/menu"
	"fmt"
)

func main() {
	Menu := menu.Menu{}

	for {
		choice := Menu.Display()
		strategy := Menu.SelectStrategy(choice)

		if strategy != nil {
			strategy.Execute()
		} else {
			fmt.Println("无效选项")
		}
	}
}
