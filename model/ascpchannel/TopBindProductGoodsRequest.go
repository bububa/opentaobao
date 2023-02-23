package ascpchannel

// TopBindProductGoodsRequest 结构体
type TopBindProductGoodsRequest struct {
	// 二级渠道编码
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 供应商渠道产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 供应商货品id
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 供应商渠道产品skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
