package alsc

// Rewards 结构体
type Rewards struct {
	// 扩展信息
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 奖品图标
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 奖品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 奖品类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 奖品值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}
