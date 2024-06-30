package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func getCPUInfo() (string, error) {
	output, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func getGPUInfo() (string, error) {
	output, err := exec.Command("system_profiler", "SPDisplaysDataType").Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Chipset Model:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}
	return "Unknown GPU", nil
}

func getMacOSVersion() (string, error) {
	output, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func formatLine(label, value string, colorFunc func(a ...interface{}) string) string {
	return fmt.Sprintf("%s: %s", label, colorFunc(value))
}

func generateWindow(lines []string, colorFunc func(a ...interface{}) string) {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}

	maxLength += 2 // Add a small buffer

	horizontalBorder := strings.Repeat("-", maxLength)

	fmt.Println(colorFunc("+" + horizontalBorder + "+"))
	for _, line := range lines {
		fmt.Println(colorFunc("| ") + line + strings.Repeat(" ", maxLength-len(line)+7) + colorFunc(" |"))
	}
	fmt.Println(colorFunc("+" + horizontalBorder + "+"))
}

func main() {
	detail := flag.Bool("detail", false, "Exibir status detalhado")
	flag.Parse()

	info, err := host.Info()
	if err != nil {
		fmt.Println("Erro ao obter informações do sistema:", err)
		return
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Erro ao obter informações de memória:", err)
		return
	}

	cpuInfo, err := getCPUInfo()
	if err != nil {
		fmt.Println("Erro ao obter informações da CPU:", err)
		return
	}

	uptime, err := host.Uptime()
	if err != nil {
		fmt.Println("Erro ao obter informações de uptime:", err)
		return
	}

	gpuInfo, err := getGPUInfo()
	if err != nil {
		fmt.Println("Erro ao obter informações da GPU:", err)
		return
	}

	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	lines := []string{
		formatLine("OS", info.OS, cyan),
		formatLine("Platform", info.Platform, cyan),
		formatLine("Kernel Version", info.KernelVersion, cyan),
		formatLine("Uptime", fmt.Sprintf("%d hours", uptime/3600), cyan),
		formatLine("Total Memory", fmt.Sprintf("%v MB", memInfo.Total/1024/1024), cyan),
		formatLine("Used Memory", fmt.Sprintf("%v MB", memInfo.Used/1024/1024), cyan),
		formatLine("CPU", cpuInfo, cyan),
		formatLine("GPU", gpuInfo, cyan),
	}

	if *detail {
		macOSVersion, err := getMacOSVersion()
		if err != nil {
			fmt.Println("Erro ao obter versão do macOS:", err)
			return
		}

		diskUsage, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Erro ao obter uso do disco:", err)
			return
		}

		netIOCounters, err := net.IOCounters(false)
		if err != nil {
			fmt.Println("Erro ao obter informações de rede:", err)
			return
		}

		lines = append(lines, []string{
			formatLine("macOS Version", macOSVersion, cyan),
			formatLine("Disk Total", fmt.Sprintf("%v GB", diskUsage.Total/1024/1024/1024), cyan),
			formatLine("Disk Used", fmt.Sprintf("%v GB", diskUsage.Used/1024/1024/1024), cyan),
			formatLine("Network Sent", fmt.Sprintf("%v MB", netIOCounters[0].BytesSent/1024/1024), cyan),
			formatLine("Network Received", fmt.Sprintf("%v MB", netIOCounters[0].BytesRecv/1024/1024), cyan),
		}...)
	}

	generateWindow(lines, green)
}
