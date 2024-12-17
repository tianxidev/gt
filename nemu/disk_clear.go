package nemu

import (
	"strconv"
	"strings"
	"time"

	"github.com/tianxidev/gt/cas"
)

// DiskClear 清理磁盘
func (n *NemuDriver) DiskClear(index string, pid int) {
	driver := n.InstallDir + `\shell\MuMuPlayerCleaner.exe`
	args := []string{"-v", "MuMuPlayer-12.0-" + index, "-p", strconv.Itoa(pid), "-isBatch", "1"}

	// clear disk
	err := new(cas.Win).ExecuteHIDE("open", driver, strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
}
