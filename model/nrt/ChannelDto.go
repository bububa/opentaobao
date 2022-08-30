package nrt

// ChannelDto 结构体
type ChannelDto struct {
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道名称
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
}
