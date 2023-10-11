package miniapp

// SellerChannelConfigDto 结构体
type SellerChannelConfigDto struct {
	// 扩展属性
	ExtProperties string `json:"ext_properties,omitempty" xml:"ext_properties,omitempty"`
	// 渠道id：1-微信；2-抖音；3-今日头条；4-快手；5-百度；6-其他
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 节点小程序id
	MiniappId int64 `json:"miniapp_id,omitempty" xml:"miniapp_id,omitempty"`
	// 0表示未配置，1表示已配置
	ConfigStatus int64 `json:"config_status,omitempty" xml:"config_status,omitempty"`
	// 是否添加到私域工作。0表示未添加，1表示添加
	AddStatus int64 `json:"add_status,omitempty" xml:"add_status,omitempty"`
}
