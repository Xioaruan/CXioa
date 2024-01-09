package createcode

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildAminExLin() {
	pwd := PwdStr()
	err := os.Chdir(pwd)
	if err != nil {
		fmt.Println("无法切换工作目录:", err)
		return
	}
	pwd2 := pwd + "\\main.go"

	buildAdminLin(pwd, pwd2)
}
func BuildExLin() {
	pwd := PwdStrAgent()
	err := os.Chdir(pwd)
	if err != nil {
		fmt.Println("无法切换工作目录:", err)
		return
	}
	pwd2 := pwd + "\\main.go"
	buildAgentLin(pwd, pwd2)
}

func buildAdminLin(pwd string, pwd2 string) {
	os.Chdir(pwd)

	shell0 := exec.Command("go", "env", "-w", "CGO_ENABLED=0", "GOOS=linux", "GOARCH=amd64") //设定为linux生成环境
	_, _ = shell0.CombinedOutput()
	shell := exec.Command("go", "build", "-o", "admin", pwd2)
	_, err := shell.CombinedOutput()
	shell1 := exec.Command("go", "env", "-w", "CGO_ENABLED=0", "GOOS=windows", "GOARCH=amd64") //修改回去
	_, _ = shell1.CombinedOutput()
	if err != nil {
		fmt.Println("build 错误", err)
	}
	nowpwd, _ := os.Getwd()

	filepath2 := filepath.Dir(nowpwd)
	newExe := filepath2 + "\\admin"
	exefile := pwd + "\\admin"
	moveFile(exefile, newExe)
}
func buildAgentLin(pwd string, pwd2 string) {

	shell0 := exec.Command("go", "env", "-w", "CGO_ENABLED=0", "GOOS=linux", "GOARCH=amd64")
	_, _ = shell0.CombinedOutput()
	shell := exec.Command("go", "build", "-o", "agent", pwd2)
	_, err := shell.CombinedOutput()
	shell1 := exec.Command("go", "env", "-w", "CGO_ENABLED=0", "GOOS=windows", "GOARCH=amd64") //修改回去
	_, _ = shell1.CombinedOutput()
	if err != nil {
		fmt.Println("build 错误", err)
	}
	nowpwd, _ := os.Getwd()

	filepath2 := filepath.Dir(nowpwd)
	newExe := filepath2 + "\\agent"
	exefile := pwd + "\\agent"
	moveFile(exefile, newExe)
}
