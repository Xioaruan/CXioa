package upload

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
	fmt.Println(bytes)
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
