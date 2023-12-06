package cpuinfo

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func GetLinuxCpuSN() string {
	// 执行dmidecode -t 4 |grep ID |sort -u |awk -F':' '{print $2}'命令
	cmd := exec.Command("bash", "-c", "/usr/sbin/dmidecode -t 4 |grep ID |sort -u |awk -F':' '{print $2}'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sndata := ""
	lines := strings.Split(string(output), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) > 0 {
			sndata = sndata + line + "\n"
		}
	}
	return sndata
}

func GetWinCpuSN() string {
	// wmic cpu get ProcessorId
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorId")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sndata := ""
	lines := strings.Split(string(output), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.Contains(line, "ProcessorId") {
			continue
		}
		sndata = sndata + line + "\n"

	}
	return sndata
}

func GetMacOSCpuSN() string {
	// system_profiler SPHardwareDataType
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sndata := ""
	lines := strings.Split(string(output), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.Trim(lines[i], " ")
		if strings.Contains(line, "Serial Number") {
			sndata = sndata + line + "\n"
		}
	}
	return sndata
}

func GetCpuSN() string {
	osType := runtime.GOOS
	sndata := ""
	if osType == "linux" {
		sndata = GetLinuxCpuSN()
	} else if osType == "windows" {
		sndata = GetWinCpuSN()
	} else if osType == "darwin" {
		sndata = GetMacOSCpuSN()
	}
	return sndata
}
