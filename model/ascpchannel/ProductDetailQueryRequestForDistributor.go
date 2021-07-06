package ascpchannel

// ProductDetailQueryRequestForDistributor 结构体
type ProductDetailQueryRequestForDistributor struct {
	// 产品 id
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 是否查询sku信息
	IncludeSku bool `json:"include_sku,omitempty" xml:"include_sku,omitempty"`
}
