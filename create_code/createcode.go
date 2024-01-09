package createcode

import (
	"fmt"
	"os"
)

func WriteIntoFile(filename string, content string) {
	filename2 := "\\" + filename
	pwd := PwdStrAgent()
	dir := pwd + filename2

	_, err := os.Stat(dir)

	if err == nil {
		err = os.Remove(dir)
		if err != nil {
			fmt.Println("无法删除文件:", err)
			return
		}
	}
	file, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	_, err = file.WriteString(content)
	// err = os.WriteFile(filename2, []byte(content), 0644)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

}
func WriteIntoAminFile(filename string, content string) {
	filename2 := "\\" + filename
	pwd := PwdStr()
	dir := pwd + filename2

	_, err := os.Stat(dir)

	if err == nil {
		err = os.Remove(dir)
		if err != nil {
			fmt.Println("无法删除文件:", err)
			return
		}
	}
	file, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	_, err = file.WriteString(content)
	// err = os.WriteFile(filename2, []byte(content), 0644)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

}
