package cpuinfo

import (
	"fmt"
	"os/exec"
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
		line := strings.Trim(lines[i], " ")
		if len(line) > 0 {
			sndata = sndata + line + "\n"
		}
	}
	return sndata
}
