package createcode

import (
	"CS2/code"
	"strings"
)

func Cr(ip string) {
	codeC := strings.Replace(code.ConnS, "%v", ip, 1)
	WriteIntoAminFile("main.go", code.MainS)
	WriteIntoAminFile("go.mod", code.Gomod)
	WriteIntoAminFile("go.sum", code.Gosum)
	WriteIntoAminFile("action\\conn.go", codeC)
	WriteIntoAminFile("upload\\upload.go", code.Upload_admin_win_upload)
	WriteIntoAminFile("action\\switch.go", code.SwtichS)
	WriteIntoAminFile("function\\conn_function.go", code.Connfunc)
	WriteIntoAminFile("function\\exec.go", code.ExecS)
	WriteIntoAminFile("function\\usefull.go", code.Use)
}
func CrAgent(ip string, port string) { //agent的写入代码
	codeC := strings.Replace(code.Func_conn, "%v", ip, 1)
	codeC = strings.Replace(codeC, "%v", port, 1)

	WriteIntoFile("main.go", code.Main)
	WriteIntoFile("go.mod", code.GomodAgent)
	WriteIntoFile("go.sum", code.GosumAgent)
	WriteIntoFile("function\\conn_function.go", codeC)
	WriteIntoFile("function\\exec.go", code.Fun_exec)
	WriteIntoFile("upload\\upload.go", code.Upload_agent_win_upload)
	WriteIntoFile("function\\usefull.go", code.Fun_use)
}
func CrAgentLinux(ip string, port string) { //agent linux版本的写入代码
	codeC := strings.Replace(code.Func_conn, "%v", ip, 1)
	codeC = strings.Replace(codeC, "%v", port, 1)

	WriteIntoFile("main.go", code.Main)
	WriteIntoFile("go.mod", code.GomodAgent)
	WriteIntoFile("go.sum", code.GosumAgent)
	WriteIntoFile("function\\conn_function.go", codeC)
	WriteIntoFile("function\\exec.go", code.AgentexecLinux)
	WriteIntoFile("upload\\upload.go", code.Upload_agent_win_upload)
	WriteIntoFile("function\\usefull.go", code.Fun_use)
}
