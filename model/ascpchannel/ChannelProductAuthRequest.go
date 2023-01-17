package ascpchannel

// ChannelProductAuthRequest 结构体
type ChannelProductAuthRequest struct {
	// 分销商名称，指定用户传值
	DistributorNick string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 授权纬度，5：经营模式 + 指定用户/6：经营模式 + 指定渠道
	DimensionType int64 `json:"dimension_type,omitempty" xml:"dimension_type,omitempty"`
	// 渠道产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
