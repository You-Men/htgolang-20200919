package manager

import (
	"fmt"

	"github.com/imsilence/user/utils"
)

// 路由
var routers = map[string]func(){}

func Register(op string, callback func()) {
	if _, ok := routers[op]; ok {
		panic(fmt.Sprintf("指令%s已经存在", op))
	}
	routers[op] = callback
}

func Run() {
	for {
		text := utils.Input("请输入指令: ")
		if text == "exit" {
			fmt.Println("退出")
			break
		}
		if action, ok := routers[text]; ok {
			action()
		} else {
			fmt.Println("指令错误")
		}
	}
}
