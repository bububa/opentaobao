package ascpchannel

// ProductDetailQueryRequestForSupplier 结构体
type ProductDetailQueryRequestForSupplier struct {
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 渠道产品 id
	ChannelProductId int64 `json:"channel_product_id,omitempty" xml:"channel_product_id,omitempty"`
	// 是否查询 sku 信息
	IncludeSku bool `json:"include_sku,omitempty" xml:"include_sku,omitempty"`
}
