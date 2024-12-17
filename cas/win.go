package cas

import (
	"syscall"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

type Win struct {
}

func New() *Win {
	return &Win{}
}

func (w *Win) Execute(lpOperation, lpFile, lpParameters string) error {
	var param *uint16
	if lpParameters != "" {
		param = syscall.StringToUTF16Ptr(lpParameters)
	}
	err := windows.ShellExecute(
		0,
		syscall.StringToUTF16Ptr(lpOperation),
		syscall.StringToUTF16Ptr(lpFile),
		param,
		nil,
		int32(win.SW_SHOW))
	return err
}

func (w *Win) ExecuteHIDE(lpOperation, lpFile, lpParameters string) error {
	var param *uint16
	if lpParameters != "" {
		param = syscall.StringToUTF16Ptr(lpParameters)
	}
	err := windows.ShellExecute(
		0,
		syscall.StringToUTF16Ptr(lpOperation),
		syscall.StringToUTF16Ptr(lpFile),
		param,
		nil,
		int32(win.SW_HIDE))
	return err
}
