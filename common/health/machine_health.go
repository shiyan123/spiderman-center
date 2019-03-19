package health

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type MachineHealth struct {
	MemTotal int
	MemFree  int
	MemUsed  int
	Usage    float64
	CPUUsed  float64
}

func GetMachineHealth() (*MachineHealth, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	cc, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	return &MachineHealth{
		MemTotal: int(v.Total / 1024 / 1024),
		MemFree:  int(v.Available / 1024 / 1024),
		MemUsed:  int(v.Used / 1024 / 1024),
		Usage:    v.UsedPercent,
		CPUUsed:  cc[0],
	}, nil
}
