package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {
	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hostname: %s\n", hostname)

	//获取操作系统信息
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("Version: %s\n", os.Getenv("OSVERSION"))

	// 获取当前用户
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %s\n", currentUser.Username)

	//获取安装的软件
	cmd := exec.Command("powershell", "Get-ItemProperty HKLM:\\Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall\\* | Select-Object DisplayName")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Installed software:\n%s\n", output)

	//获取代理配置
	cmd2 := exec.Command("powershell", "Get-ItemProperty HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings | Select-Object ProxyServer")
	output2, err2 := cmd2.Output()
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("Proxy configuration:\n%s\n", output2)

	//获取PuTTY对话和密钥
	currentUser3, err3 := user.Current()
	if err3 != nil {
		panic(err3)
	}

	puttySessionDir := filepath.Join(currentUser3.HomeDir, "AppData", "Roaming", "PuTTY")
	fmt.Printf("PuTTY sessions directory: %s\n", puttySessionDir)

	privateKeyDir := filepath.Join(currentUser3.HomeDir, ".ssh")
	fmt.Printf("SSH private key directory: %s\n", privateKeyDir)

	//获取过去的RDP会议
	cmd4 := exec.Command("powershell", "Get-EventLog -LogName Security | Where-Object {$_.EventID -eq 4624} | Select-Object TimeGenerated, Message")
	output4, err4 := cmd4.Output()
	if err4 != nil {
		panic(err4)
	}
	fmt.Printf("RDP sessions:\n%s\n", output4)

	// 获取过去的运行命令
	cmd5 := exec.Command("powershell", "Get-EventLog -LogName 'Windows PowerShell' | Where-Object {$_.EventID -eq 400} | Format-Table TimeGenerated, Message")
	out5, err5 := cmd5.Output()
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println(string(out5))
}
