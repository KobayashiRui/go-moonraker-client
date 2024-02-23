package moonrakerclient

import (
	"encoding/json"
	"fmt"
)

func GetNotifyStatusUpdate(params *json.RawMessage) map[string]interface{} {
	var get_data []interface{}
	err := json.Unmarshal(*params, &get_data)
	if err != nil {
		fmt.Println("ERROR notify_proc_stat_update")
	}
	return get_data[0].(map[string]interface{})
}

type PrinterObjectsQuery struct {
	Eventtime float32                `json:"eventtime"`
	Status    map[string]interface{} `json:"status"`
}

type IdleTimeout struct {
	PrintingTime float32 `json:"printing_time"`
	State        string  `json:"state"`
}

type VirtualSdcard struct {
	Progress     float32 `json:"progress"`
	IsActive     bool    `json:"is_active"`
	FilePosition int     `json:"file_position"`
}

type PrintStatsInfo struct {
	TotalLayer   *int32 `json:"total_layer"`
	CurrentLayer *int32 `json:"current_layer"`
}

type PrintStats struct {
	Filename      string         `json:"filename"`
	TotalDuration float32        `json:"total_duration"`
	PrintDuration float32        `json:"print_duration"`
	FilamentUsed  float32        `json:"filament_used"`
	State         string         `json:"state"`
	Message       string         `json:"message"`
	Info          PrintStatsInfo `json:"info"`
}

type DisplayStatus struct {
	Message  string  `json:"message"`
	Progress float32 `json:"progress"`
}

func ConvRef(bd interface{}, res interface{}) {
	// Convert map to string
	jsonStr, err := json.Marshal(bd)
	if err != nil {
		fmt.Println(err)
	}

	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, res); err != nil {
		fmt.Println(err)
	}
}

func Conv[T any](bd interface{}) T {
	// Convert map to string
	jsonStr, err := json.Marshal(bd)
	if err != nil {
		fmt.Println(err)
	}

	// Convert json string to struct
	var res T
	if err := json.Unmarshal(jsonStr, &res); err != nil {
		fmt.Println(err)
	}
	return res
}
