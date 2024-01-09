package createcode

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPwdS() {
	pwd, _ := os.Getwd()
	pwd += "\\CSadmin"

	createAction(pwd)
	createfunction(pwd)
	createUpload(pwd)
}
func GetPwdAgent() {
	pwd, _ := os.Getwd()

	uppwd := filepath.Dir(pwd)
	os.Chdir(uppwd)
	pwd, _ = os.Getwd()

	pwd += "\\CSagent"

	createAction(pwd)
	createfunction(pwd)
	createUpload(pwd)
}

func createAction(pwd string) {
	newF := pwd + "\\action"

	err := os.MkdirAll(newF, 0755)
	if err != nil {
		fmt.Println("无法创建文件夹:", err)
		return
	}
}
func createUpload(pwd string) {
	newF := pwd + "\\upload"

	err := os.MkdirAll(newF, 0755)
	if err != nil {
		fmt.Println("无法创建文件夹:", err)
		return
	}
}
func createfunction(pwd string) {
	newF := pwd + "\\function"

	err := os.MkdirAll(newF, 0755)
	if err != nil {
		fmt.Println("无法创建文件夹:", err)
		return
	}
}

func PwdStr() string {
	pwd, _ := os.Getwd()
	pwd += "\\CSadmin"
	return pwd
}
func PwdStrAgent() string {
	pwd, _ := os.Getwd()
	pwd += "\\CSagent"

	return pwd
}
