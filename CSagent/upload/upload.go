package upload

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

