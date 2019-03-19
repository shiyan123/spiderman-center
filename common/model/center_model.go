package model

import (
	"spiderman-center/common/health"
)

type Node struct {
	ID          string
	IP          string
	Name        string
	MachineInfo *health.MachineHealth
	TaskMap     map[string]*TaskInfo
}
