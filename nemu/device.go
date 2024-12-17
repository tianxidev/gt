package nemu

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
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
