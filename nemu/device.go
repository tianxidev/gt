package nemu

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/tianxidev/gt/cas"
)

// connect 连接设备
func (n *NemuDriver) connect(index int) {
	ptrPath := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(n.InstallDir)))
	ret, _, err := n.NemuConnect.Call(ptrPath, uintptr(index))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connect: ", ret)
	n.ConnectionId = int(ret)
}

// disconnect 断开连接
func (n *NemuDriver) disconnect() {
	ret, _, err := n.NemuDisconnect.Call(uintptr(n.ConnectionId))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("disconnect", ret)
}

// StartDevice 启动设备
func (n *NemuDriver) StartDevice(index int) {
	n.connect(index)
	n.disconnect()
}

// GetDeviceStatus 查询单个模拟器状态信息
func (n *NemuDriver) GetDeviceStatus(index int) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"info", "-v", strconv.Itoa(index)}

	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// GetDevicesStatus 查询所有模拟器状态信息
func (n *NemuDriver) GetDevicesStatus() (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"info", "-v", "all"}

	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// CreateDevice 创建一个新的模拟器
func (n *NemuDriver) CreateDevice() (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"create"}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// CreateDevices 批量创建模拟器
func (n *NemuDriver) CreateDevices(count int) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"create", "-n", strconv.Itoa(count)}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// CopyDevice 复制模拟器
func (n *NemuDriver) CopyDevice(index int) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"clone", "-v", strconv.Itoa(index)}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// CopyDevices 批量复制模拟器
func (n *NemuDriver) CopyDevices(index int, count int) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"clone", "-v", strconv.Itoa(index), "-n", strconv.Itoa(count)}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// DeleteDevice 删除模拟器
func (n *NemuDriver) DeleteDevice(index int) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"delete", "-v", strconv.Itoa(index)}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// 重命名模拟器
func (n *NemuDriver) RenameDevice(index int, name string) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"rename", "-v", strconv.Itoa(index), "-n", name}
	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// StartDeviceByIndex 启动指定模拟器
func (n *NemuDriver) StartDeviceByIndex(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "launch"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// 启动所有模拟器
func (n *NemuDriver) StartAllDevices() {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", "all", "launch"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// StopDeviceByIndex 停止指定模拟器
func (n *NemuDriver) StopDeviceByIndex(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "shutdown"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// StopAllDevices 停止所有模拟器
func (n *NemuDriver) StopAllDevices() {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", "all", "shutdown"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// RestartDeviceByIndex 重启指定模拟器
func (n *NemuDriver) RestartDeviceByIndex(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "restart"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// 模拟器窗口排版
func (n *NemuDriver) ArrangeDevices() {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"sort"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// StartApp 启动模拟器里的应用
func (n *NemuDriver) StartApp(index int, packageName string) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "app", "launch", "-pkg", packageName}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// 在所有模拟器上启动应用
func (n *NemuDriver) StartAppOnAllDevices(packageName string) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", "all", "app", "launch", "-pkg", packageName}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// StopApp 停止模拟器里的应用
func (n *NemuDriver) StopApp(index int, packageName string) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "app", "close", "-pkg", packageName}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// StopAppOnAllDevices 在所有模拟器上停止应用
func (n *NemuDriver) StopAppOnAllDevices(packageName string) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", "all", "app", "close", "-pkg", packageName}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// GetAppInfo 指定模拟器上查询应用信息
func (n *NemuDriver) GetAppInfo(index int, packageName string) (string, error) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "app", "info", "-pkg", packageName}

	out, err := new(cas.Win).ExecuteWithOutput(driver, strings.Join(args, " "))
	if err != nil {
		return "", err
	}
	return out, nil
}

// RotateScreen 模拟器屏幕旋转
func (n *NemuDriver) RotateScreen(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "tool", "func", "-n", "rotate"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// BackToHome 模拟器返回主页
func (n *NemuDriver) BackToHome(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "tool", "func", "-n", "go_home"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// GoBack 模拟器返回操作
func (n *NemuDriver) GoBack(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "tool", "func", "-n", "go_back"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}

// Shake 模拟器摇一摇
func (n *NemuDriver) Shake(index int) {
	driver := n.InstallDir + `\shell\MuMuManager.exe`
	args := []string{"control", "-v", strconv.Itoa(index), "tool", "func", "-n", "shake"}

	err := new(cas.Win).Execute("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
}
