package nemu

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

type NemuDriver struct {
	InstallDir     string
	NemuConnect    *syscall.LazyProc
	NemuDisconnect *syscall.LazyProc
	ConnectionId   int
}

func New() *NemuDriver {
	regSrc := `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\MuMuPlayer-12.0`

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, regSrc, registry.READ|registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	DisplayVersion, _, err := key.GetStringValue("DisplayVersion")
	if err != nil {
		log.Fatal(err)
	}

	if strings.Compare(DisplayVersion, "4.1") < 0 {
		log.Fatal("nemu version is too low, please upgrade to 4.1 or above")
	}

	UninstallString, _, err := key.GetStringValue("UninstallString")
	if err != nil {
		log.Fatal(err)
	}

	lastIndex := strings.LastIndex(UninstallString, `\`)
	if lastIndex == -1 {
		log.Fatal("UninstallString is invalid")
	}

	installDir := UninstallString[1:lastIndex]

	dllSrc := installDir + `\shell\sdk\external_renderer_ipc.dll`
	dll := syscall.NewLazyDLL(dllSrc)

	NemuConnect := dll.NewProc("nemu_connect")
	NemuDisconnect := dll.NewProc("nemu_disconnect")

	return &NemuDriver{
		InstallDir:     installDir,
		NemuConnect:    NemuConnect,
		NemuDisconnect: NemuDisconnect,
	}
}

func (n *NemuDriver) connect(index int) {
	ptrPath := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(n.InstallDir)))
	ret, _, err := n.NemuConnect.Call(ptrPath, uintptr(index))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("connect: ", ret)
	n.ConnectionId = int(ret)
}

func (n *NemuDriver) disconnect() {
	ret, _, err := n.NemuDisconnect.Call(uintptr(n.ConnectionId))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("disconnect", ret)
}

func (n *NemuDriver) StartDevice(index int) {
	n.connect(index)

	n.disconnect()
}
