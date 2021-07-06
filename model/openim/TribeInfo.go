package openim

// TribeInfo 结构体
type TribeInfo struct {
	// 群头像URL地址
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 群名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 群公告
	Notice string `json:"notice,omitempty" xml:"notice,omitempty"`
	// 群ID
	TribeId int64 `json:"tribe_id,omitempty" xml:"tribe_id,omitempty"`
	// 群验证模式
	CheckMode int64 `json:"check_mode,omitempty" xml:"check_mode,omitempty"`
	// 群类型
	TribeType int64 `json:"tribe_type,omitempty" xml:"tribe_type,omitempty"`
	// 群接收标记
	RecvFlag int64 `json:"recv_flag,omitempty" xml:"recv_flag,omitempty"`
}
