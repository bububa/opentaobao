package trade

// ChannelOrderItem 
/* model for simplify = false
type ChannelOrderItem struct {

    // 商品名称
    
    ItemName   string `json:"item_name,omitempty"`
    

    // 分销价格(分)
    
    DistributionPrice   int64 `json:"distribution_price,omitempty"`
    

    // 条形码
    
    Barcode   string `json:"barcode,omitempty"`
    

    // skuId
    
    SkuId   string `json:"sku_id,omitempty"`
    

    // 商品id（商品id和货号必填一个）
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 货号（商品id和货号必填一个）
    
    InventoryNo   string `json:"inventory_no,omitempty"`
    

    // 数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

}
*/

// ChannelOrderItem 
type ChannelOrderItem struct {

    // 商品名称
    ItemName   string `json:"item_name,omitempty"`

    // 分销价格(分)
    DistributionPrice   int64 `json:"distribution_price,omitempty"`

    // 条形码
    Barcode   string `json:"barcode,omitempty"`

    // skuId
    SkuId   string `json:"sku_id,omitempty"`

    // 商品id（商品id和货号必填一个）
    ItemId   string `json:"item_id,omitempty"`

    // 货号（商品id和货号必填一个）
    InventoryNo   string `json:"inventory_no,omitempty"`

    // 数量
    Quantity   int64 `json:"quantity,omitempty"`

}
