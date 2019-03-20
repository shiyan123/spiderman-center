package model

import (
	"spiderman-center/common/health"
)

type Node struct {
	ID          string                `json:"id"`
	IP          string                `json:"ip"`
	Name        string                `json:"name"`
	MachineInfo *health.MachineHealth `json:"machineInfo"`
	TaskMap     map[string]*TaskInfo  `json:"taskMap"`
}
