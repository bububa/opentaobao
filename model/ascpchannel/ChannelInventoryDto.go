package ascpchannel

// ChannelInventoryDto 
type ChannelInventoryDto struct {
    // 库存数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 产品ID
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 二级渠道
    SubChannelCode   string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
    // skuId
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 市场
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}
