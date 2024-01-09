package createcode

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildAminEx() {
	pwd := PwdStr()
	err := os.Chdir(pwd)
	if err != nil {
		fmt.Println("无法切换工作目录:", err)
		return
	}
	pwd2 := pwd + "\\main.go"

	buildAdmin(pwd, pwd2)
}
func BuildEx() {
	pwd := PwdStrAgent()
	err := os.Chdir(pwd)

	if err != nil {
		fmt.Println("无法切换工作目录:", err)
		return
	}
	pwd2 := pwd + "\\main.go"

	buildAgent(pwd, pwd2)


}

func moveFile(exe string, pwd string) {
	err := os.Rename(exe, pwd)
	if err != nil {
		fmt.Println("无法移动文件:", err)
		return
	}
}
func buildAdmin(pwd string, pwd2 string) {
	os.Chdir(pwd)


	shell := exec.Command("go", "build", "-o", "admin.exe", pwd2)
	_, err := shell.CombinedOutput()
	if err != nil {
		fmt.Println("build 错误", err)
	}
	nowpwd, _ := os.Getwd()


	filepath2 := filepath.Dir(nowpwd)
	newExe := filepath2 + "\\admin.exe"
	exefile := pwd + "\\admin.exe"
	moveFile(exefile, newExe)
}
func buildAgent(pwd string, pwd2 string) {
	// pwd4, _ := os.Getwd()
	// fmt.Println(pwd4)
	shell := exec.Command("go", "build", "-o", "agent.exe", "-ldflags", "-H=windowsgui", pwd2)
	_, err := shell.CombinedOutput()
	if err != nil {
		fmt.Println("build 错误", err)
	}
	nowpwd, _ := os.Getwd()


	filepath2 := filepath.Dir(nowpwd)
	newExe := filepath2 + "\\agent.exe"
	exefile := pwd + "\\agent.exe"
	moveFile(exefile, newExe)
}
