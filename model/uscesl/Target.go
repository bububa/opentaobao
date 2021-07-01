package uscesl

// Target 结构体
type Target struct {
	// AP的mac地址
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// 是否在线
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 是否激活
	IsActivate bool `json:"is_activate,omitempty" xml:"is_activate,omitempty"`
	// 型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
