package alsc

// MissionExt 结构体
type MissionExt struct {
	// 其他信息
	ExtValue string `json:"ext_value,omitempty" xml:"ext_value,omitempty"`
	// 扩展类型
	ExtType string `json:"ext_type,omitempty" xml:"ext_type,omitempty"`
}
