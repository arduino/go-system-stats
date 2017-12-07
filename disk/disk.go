package disk

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
)

// FSStats contains usage stats for a filesystem
type FSStats struct {
	Device         string
	Type           string
	MountPoint     string
	FreeSpace      uint64
	AvailableSpace uint64
	DiskSize       uint64
}

// GetStats returns usage stats for all mounted filesystems
func GetStats() ([]*FSStats, error) {
	data, err := ioutil.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return nil, fmt.Errorf("detecting mounted filesystems: %s", err)
	}
	mountInfos := strings.Split(string(data), "\n")

	res := []*FSStats{}
	for _, mountInfo := range mountInfos {
		fields := strings.Fields(mountInfo)
		if len(fields) < 5 {
			continue
		}

		mount := fields[4]

		_, err := os.Stat(mount)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("getting stat for %s: %s", mount, err)
		}

		statfs := &syscall.Statfs_t{}
		if err := syscall.Statfs(mount, statfs); err != nil {
			// ignore error
			continue
		}

		fs := &FSStats{
			MountPoint:     mount,
			Type:           fields[8],
			Device:         fields[9],
			FreeSpace:      statfs.Bfree * uint64(statfs.Bsize),
			AvailableSpace: statfs.Bavail * uint64(statfs.Bsize),
			DiskSize:       statfs.Blocks * uint64(statfs.Bsize),
		}
		res = append(res, fs)
	}

	return res, nil
}
