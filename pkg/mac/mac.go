package mac

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GetLinuxMac() string {
	// 执行ip link show命令
	cmd := exec.Command("ip", "link", "show")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// fmt.Println(string(output))
	macdata := ""
	// 从输出结果中解析MAC地址
	lines := strings.Split(string(output), "\n")
	for i := 0; i < len(lines); i += 2 {
		line := lines[i]
		// 定义正则表达式
		re := regexp.MustCompile(`\d+:\s+(\w+):`)
		match := re.FindStringSubmatch(line)
		interfaceName := ""
		if len(match) > 1 {
			interfaceName = match[1]
		}
		if strings.HasPrefix(interfaceName, "eth") || strings.HasPrefix(interfaceName, "en") || strings.HasPrefix(interfaceName, "wl") {
			lineNext := lines[i+1]
			if strings.Contains(lineNext, "link/ether") {
				startIndex := strings.Index(lineNext, "link/ether") + 11
				endIndex := startIndex + 17
				mac := lineNext[startIndex:endIndex]
				macdata = macdata + mac + "\n"
			}
		}
	}
	return macdata
}

func GetWinMac() string {

	return ""
}
