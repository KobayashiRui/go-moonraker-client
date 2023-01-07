package moonrakerclient

type MoonrakerStats struct {
	Time     float64 `json:"time"`
	CpuUsage float32 `json:"cpu_usage"`
	Memory   int32   `json:"memory"`
	MemUnit  string  `json:"mem_units"`
}

type Network struct {
	RxBytes   int32   `json:"rx_bytes"`
	TxBytes   int32   `json:"tx_bytes"`
	BandWidth float32 `json:"bandwidth"`
}

type NotifyProcStatUpdate struct {
	MoonrakerStats       MoonrakerStats     `json:"moonraker_stats"`
	CpuTemp              float32            `json:"cpu_temp"`
	Network              map[string]Network `json:"network"`
	SystemCpuUsage       map[string]float32 `json:"system_cpu_usage"`
	WebsocketConnections int32              `json:"websocket_connections"`
}

type PrinterObjects struct {
	Objects []string `json:"objects"`
}

type GcodeFileData struct {
	Path        string  `json:"path"`
	Modified    float64 `json:"modified"`
	Size        int64   `json:"size"`
	Permissions string  `json:"permissions"`
}
