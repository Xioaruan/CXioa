package action

import (
	"CSadmin/function"
	"fmt"
)

func SwitchChoose() string {
	fmt.Println("【1】 执行命令 【2】 文件上传 【3】 暂未开发")
	fmt.Print("choose >")
	var choose string

	function.Scanf(&choose)
	return choose
}

func FileChoose() string {

	fmt.Println("【1】 上传文件 【2】 下载文件 ")
	fmt.Print("choose >")
	var choose string

	function.Scanf(&choose)
	return choose
}
