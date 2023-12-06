package mac

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
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

func isVirtualMAC(macAddress string) bool {
	virtualKeywords := []string{"Virtual", "VMware", "VirtualBox"}

	for _, keyword := range virtualKeywords {
		if strings.Contains(macAddress, keyword) {
			return true
		}
	}
	return false
}

func parseMACAddresses(output string) []string {
	var macAddresses []string

	lines := strings.Split(output, "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) <= 0 {
			continue
		}
		// fmt.Println(line)
		if strings.HasPrefix(line, "Description") || strings.HasPrefix(line, "描述") {
			if !isVirtualMAC(line) {
				i += 2
				continue
			}
		}
		if strings.HasPrefix(line, "Physical Address") || strings.HasPrefix(line, "物理地址") {
			// fields := strings.Fields(line)
			lineItem := strings.Split(line, ":")
			if len(lineItem) > 1 {
				macAddress := strings.TrimSpace(lineItem[1])
				macAddresses = append(macAddresses, macAddress)
			}
		}
	}

	return macAddresses
}

func GetWinMac() string {
	// cmd := exec.Command("ipconfig", "/all")
	cmd := exec.Command("cmd", "/c", "chcp 65001 & ipconfig /all")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	macdata := ""
	macAddresses := parseMACAddresses(string(output))
	for _, mac := range macAddresses {
		macOk := strings.TrimSpace(mac)
		if len(macOk) > 0 {
			macdata = macdata + macOk + "\n"
		}
	}
	return macdata
}

func GetMacOSMac() string {

	return ""
}

func GetMacAddr() string {
	osType := runtime.GOOS
	macdata := ""
	if osType == "linux" {
		macdata = GetLinuxMac()
	} else if osType == "windows" {
		macdata = GetWinMac()
	} else if osType == "darwin" {
		macdata = GetMacOSMac()
	}
	return macdata
}
