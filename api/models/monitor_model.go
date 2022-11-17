package models

import "time"

type NetworkDevice struct {
	MachineID int `json:"machineId"`
	Status
	LastLoggedIn string `json:"lastLoggedIn"`
	SysTime time.Time `json:"sysTime"`
}

type Status struct {
	CpuTemp int `json:"cpuTemp"`
	FanSpeed int `json:"fanSpeed"`
	HddSpace int `json:"HDDSpace"`
}