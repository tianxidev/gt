package nemu

import (
	"log"
	"strings"
	"syscall"

	"github.com/lxn/win"
	"golang.org/x/sys/windows/registry"
)

type NemuDriver struct {
	InstallDir     string
	NemuConnect    *syscall.LazyProc
	NemuDisconnect *syscall.LazyProc
	ConnectionId   int
}

func New() *NemuDriver {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)

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
