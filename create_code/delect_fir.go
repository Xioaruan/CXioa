package createcode

import (
	"fmt"
	"os"
	"path/filepath"
)

func DelectFIR() {
	pwd, _ := os.Getwd()

	delectAction(pwd)
}

func delectAction(pwd string) {

	pwd2, _ := os.Getwd()
	filemain := filepath.Dir(pwd2)
	os.Chdir(filemain)

	err := os.RemoveAll(pwd)
	if err != nil {
		fmt.Println("无法删除文件夹:", err)
		return
	}
	pwd3 := filemain + "\\CSadmin"

	err = os.RemoveAll(pwd3)
	if err != nil {
		fmt.Println("无法删除文件夹:", err)
		return
	}
}

func DelectFIRLin() {
	pwd, _ := os.Getwd()

	delectActionLin(pwd)
}

func delectActionLin(pwd string) {

	pwd2, _ := os.Getwd()
	filemain := filepath.Dir(pwd2)
	os.Chdir(filemain)
	err := os.RemoveAll(pwd)
	if err != nil {
		fmt.Println("无法删除文件夹:", err)
		return
	}
	pwd3 := filemain + "\\CSadmin"

	err = os.RemoveAll(pwd3)
	if err != nil {
		fmt.Println("无法删除文件夹:", err)
		return
	}
}
