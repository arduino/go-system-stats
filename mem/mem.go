package mem

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Stats contains data about the memory avaible/used
type Stats struct {
	FreeMem      uint64
	TotalMem     uint64
	AvailableMem uint64
	Buffers      uint64
	Cached       uint64

	FreeSwapMem  uint64
	TotalSwapMem uint64
}

// GetStats returns the current memory usage status
func GetStats() (*Stats, error) {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, fmt.Errorf("detecting mounted filesystems: %s", err)
	}
	memInfos := strings.Split(string(data), "\n")

	res := &Stats{}
	for _, l := range memInfos {
		fields := strings.Fields(l)
		if len(fields) < 2 {
			continue
		}
		tag := fields[0]
		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}
		switch tag {
		case "MemTotal:":
			res.TotalMem = val
		case "MemFree:":
			res.FreeMem = val
		case "MemAvailable:":
			res.AvailableMem = val
		case "Buffers:":
			res.Buffers = val
		case "Cached:":
			res.Cached = val
		case "SwapTotal:":
			res.TotalSwapMem = val
		case "SwapFree:":
			res.FreeSwapMem = val
		}
	}
	return res, nil
}
