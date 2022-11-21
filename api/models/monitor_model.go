package models

import (
	"github.com/google/uuid"
	"time"
)

type NetworkDevice struct {
	MachineID    uuid.UUID `json:"machineId" validate:"required"`
	Status       `json:"status" validate:"required"`
	LastLoggedIn string    `json:"lastLoggedIn" validate:"required"`
	SysTime      time.Time `json:"sysTime" validate:"required"`
}

type Status struct {
	CpuTemp  int `json:"cpuTemp" validate:"required"`
	FanSpeed int `json:"fanSpeed" validate:"required"`
	HddSpace int `json:"HDDSpace" validate:"required"`
}
