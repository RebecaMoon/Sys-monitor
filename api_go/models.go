package main

type Metric struct { // This defines the structure of the data we receive from the agent.
	ID            int     `json:"id"`
	CPUPercent    float64 `json:"cpu_percent"`    // The fields start with uppercase letters because in Go, if a field starts with an uppercase letter, it is exported.
	MemoryPercent float64 `json:"memory_percent"` // That means other packages can access it, such as the encoding/json package in this case.
	DiskPercent   float64 `json:"disk_percent"`
	Timestamp     float64 `json:"timestamp"`
}
