package miniapp

// AppChannelQueryRequest 结构体
type AppChannelQueryRequest struct {
	// 渠道id
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 小程序id
	MiniappId int64 `json:"miniapp_id,omitempty" xml:"miniapp_id,omitempty"`
}
