package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hlinfocc/hardware/pkg/cpuinfo"
	"github.com/hlinfocc/hardware/pkg/mac"
	"github.com/hlinfocc/hardware/pkg/version"
)

type Resp struct {
	Code int
	Msg  string
	Data string
}

/**
* 命令行参数结构体
 */
type Args struct {
	Version bool
	Sn      bool
}

/**
* 初始化命令行参数信息
 */
func initParams() Args {
	args := Args{}
	flag.BoolVar(&args.Version, "v", args.Version, "显示版本信息")
	flag.BoolVar(&args.Sn, "sn", args.Sn, "直接获取CPU序列号及网卡MAC地址信息")
	flag.Parse()
	return args
}

/**
* 启动Socket服务
 */
func StartServer() {
	socketPath := "/var/run/hlinfo-hardware.socket"
	os.Remove(socketPath)
	tcpAddr, err := net.ResolveUnixAddr("unix", socketPath)
	checkError(err)
	listener, err := net.ListenUnix("unix", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go HandleServerConn(conn)
	}
}

func HandleServerConn(conn net.Conn) {
	// 设置2分钟超时时间
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	// 将最大请求长度设置为128B以防止DDos攻击
	request := make([]byte, 128)
	// 退出前关闭连接
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			// 客户端已关闭连接
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "SN" {
			cpusn := cpuinfo.GetCpuSN()
			mac := mac.GetMacAddr()
			sninfo := cpusn + "" + mac
			conn.Write([]byte(sninfo))
		} else {
			fmt.Println(strings.TrimSpace(string(request[:read_len])))
			rs := Resp{}
			rs.Code = 200
			rs.Msg = "获取成功"
			rs.Data = "[{\"cpusn\":\"\"},{\"mac\":\"\"},{\"hdd\":\"\"}]"
			v, _ := json.Marshal(rs)
			conn.Write([]byte(string(v)))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func checkPortStatus(port int) bool {
	// 监听 端口
	listenerPort := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", listenerPort)
	if err != nil {
		// 如果监听失败，则说明端口已被占用
		return false
	}
	// 关闭监听器
	defer listener.Close()

	// 如果监听成功，则说明端口未被占用
	return true
}

func writePort(port int) {
	filePath := "/var/run/hlinfo-hardware.port"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(port))
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
}

func StartWebServer() {
	// 定义处理请求的函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 获取客户端传递的参数
		sn := r.FormValue("sn")

		// 根据参数进行相应的响应
		if sn != "" {
			cpusn := cpuinfo.GetCpuSN()
			mac := mac.GetMacAddr()
			sninfo := cpusn + "" + mac
			fmt.Fprintf(w, sninfo)
		} else {
			fmt.Fprintf(w, "Error")
		}
	}
	port := 1840
	for !checkPortStatus(port) {
		port = port + 1
	}
	writePort(port)
	httpPort := fmt.Sprintf(":%d", port)
	// 注册处理函数并启动 Web 服务
	http.HandleFunc("/", handler)
	http.ListenAndServe(httpPort, nil)

}

func main() {
	args := initParams()

	if args.Version {
		fmt.Println(version.Full())
	} else if args.Sn {
		cpusn := cpuinfo.GetCpuSN()
		mac := mac.GetMacAddr()
		sninfo := cpusn + "" + mac
		fmt.Println(sninfo)
	} else {
		go StartWebServer()
		StartServer()
	}
}
