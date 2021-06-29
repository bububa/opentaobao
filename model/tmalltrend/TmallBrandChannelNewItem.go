package tmalltrend

// TmallBrandChannelNewItem 
type TmallBrandChannelNewItem struct {

    // 商品sku在渠道的库存
    
    ChannelQuantity   int64 `json:"channel_quantity,omitempty" xml:"channel_quantity,omitempty"`
    

    // 渠道商品id
    
    ChannelItemId   int64 `json:"channel_item_id,omitempty" xml:"channel_item_id,omitempty"`
    

    // 天猫品牌名称
    
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    

    // 发售渠道名称
    
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    

    // 渠道商品skuId
    
    ChannelSkuId   int64 `json:"channel_sku_id,omitempty" xml:"channel_sku_id,omitempty"`
    

    // 在渠道的发售时间
    
    ChannelPublishTime   string `json:"channel_publish_time,omitempty" xml:"channel_publish_time,omitempty"`
    

    // 是否天猫同款
    
    TmallSame   bool `json:"tmall_same,omitempty" xml:"tmall_same,omitempty"`
    

    // 天猫商品id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 天猫商品名称
    
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    

    // 渠道商品名称
    
    ChannelItemName   string `json:"channel_item_name,omitempty" xml:"channel_item_name,omitempty"`
    

    // 商品sku在渠道的图片
    
    ChannelItemPicts   []string `json:"channel_item_picts,omitempty" xml:"channel_item_picts>string,omitempty"`
    

    // 商品sku在渠道对应的属性对，格式为k1:v1;k2:v2
    
    ChannelProperties   string `json:"channel_properties,omitempty" xml:"channel_properties,omitempty"`
    

    // 天猫品牌id
    
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    

    // 商品sku在渠道的发售价格
    
    ChannelItemPrice   string `json:"channel_item_price,omitempty" xml:"channel_item_price,omitempty"`
    

    // barcode码
    
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    

    // 天猫商品skuId
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 渠道的发售区域
    
    ChannelPublishArea   string `json:"channel_publish_area,omitempty" xml:"channel_publish_area,omitempty"`
    

}
