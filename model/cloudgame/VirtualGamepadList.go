package cloudgame

// VirtualGamepadList 结构体
type VirtualGamepadList struct {
	// 手柄配置
	Config string `json:"config,omitempty" xml:"config,omitempty"`
	// 配置名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 优先级
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 手柄ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 手柄类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
