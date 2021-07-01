package trade

// ChannelOrderOption 结构体
type ChannelOrderOption struct {
	// 是否允许供应商修改
	IsAllowUpperModify bool `json:"is_allow_upper_modify,omitempty" xml:"is_allow_upper_modify,omitempty"`
}
