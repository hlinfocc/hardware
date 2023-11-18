package hdd

import (
	"fmt"
	"os/exec"
)

func GetLinuxHddSN() string {
	// 执行lsblk -d -n -o serial命令
	cmd := exec.Command("bash", "-c", "lsblk -d -n -o serial")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(output)
}
