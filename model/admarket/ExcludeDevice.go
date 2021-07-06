package admarket

// ExcludeDevice 结构体
type ExcludeDevice struct {
	// 排他设备id
	DeviceIds []string `json:"device_ids,omitempty" xml:"device_ids>string,omitempty"`
	// 排他开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 排他结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}
