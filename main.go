package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/matryer/try"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func getGPUInfo() (string, error) {
	var output []byte
	err := try.Do(func(attempt int) (bool, error) {
		var err error
		output, err = exec.Command("wmic", "path", "win32_videocontroller", "get", "name").Output()
		return attempt < 3, err
	})
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if trimmedLine := strings.TrimSpace(line); trimmedLine != "" && !strings.Contains(trimmedLine, "Name") {
			return trimmedLine, nil
		}
	}
	return "Unknown GPU", nil
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

	horizontalBorder := strings.Repeat("-", maxLength)

	fmt.Println(colorFunc("+" + horizontalBorder + "+"))
	for _, line := range lines {
		fmt.Println(colorFunc("| ") + line + strings.Repeat(" ", maxLength-len(line)+7) + colorFunc(" |"))
	}
	fmt.Println(colorFunc("+" + horizontalBorder + "+"))
}

func main() {
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

	cpuInfo, err := cpu.Info()
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
		formatLine("CPU", cpuInfo[0].ModelName, cyan),
		formatLine("GPU", gpuInfo, cyan),
	}

	generateWindow(lines, green)
}
