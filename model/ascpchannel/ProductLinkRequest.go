package ascpchannel

// ProductLinkRequest 
type ProductLinkRequest struct {
    // 分销商商品 ID
    OutItemId   string `json:"out_item_id,omitempty" xml:"out_item_id,omitempty"`
    // sku 列表
    SkuList   []ProductSkuLinkDTO `json:"sku_list,omitempty" xml:"sku_list>product_sku_link_dto,omitempty"`
    // 供应商产品 id
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 经营模式
    SalesMode   string `json:"sales_mode,omitempty" xml:"sales_mode,omitempty"`
    // 二级渠道编码
    SubChannelCode   string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
    // 渠道编码
    ChannelCode   string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}
