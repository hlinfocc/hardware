package cpuinfo

import (
	"fmt"
	"os/exec"
)

func GetLinuxCpuInfo() string {
	// 执行dmidecode -t 4 |grep ID |sort -u |awk -F':' '{print $2}'命令
	cmd := exec.Command("bash", "-c", "dmidecode -t 4 |grep ID |sort -u |awk -F':' '{print $2}'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(output)
}
