package ascpchannel

// ItemDistributePublishRequest 结构体
type ItemDistributePublishRequest struct {
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道产品id
	SourceProductId int64 `json:"source_product_id,omitempty" xml:"source_product_id,omitempty"`
}
