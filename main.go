package main

import (
	createcode "CS2/create_code"
	"bufio"
	"fmt"
	"os"
)

var ip string

var port string
var system string

func main() {
	fmt.Print("IP >")
	Scanf(&ip)
	fmt.Print("PORT >")
	Scanf(&port)
	fmt.Print("SYSTEM > [1] win [2] linux [3] All\n")
	fmt.Print("SYSTEM > ")
	Scanf(&system)
	fmt.Println("运行完成自动关闭，请等待。。。。。")
	if system == "1" {
		BuildWin()
		os.Exit(1)
	} else if system == "2" {
		BuildLinux()
		os.Exit(1)
	} else if system == "3" {
		BuildWin()
		BuildLinux()
		os.Exit(1)
	} else {
		fmt.Println("not ye")
	}
}
func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	*a = string(data)
}
func BuildWin() {
	buildadmin()
	buildagent()
	// createcode.DelectFIR()
}
func BuildLinux() {
	buildadminLinux() //linux版本的admin不需要改变，只有在build的时候需要改动
	buildagentlinux()
	// createcode.DelectFIRLin()
}

func buildadmin() {
	createcode.GetPwdS()
	createcode.Cr(port)
	createcode.BuildAminEx()
}
func buildagent() {
	createcode.GetPwdAgent()
	createcode.CrAgent(ip, port)
	createcode.BuildEx()
}

func buildadminLinux() {
	createcode.GetPwdS()
	createcode.Cr(port)
	createcode.BuildAminExLin() //这里需要改动
}
func buildagentlinux() {
	createcode.GetPwdAgent()
	createcode.CrAgentLinux(ip, port)
	createcode.BuildExLin()
}
