package aliexpresssumaitong

// Item 
type Item struct {

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品价格
    
    Price   float64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 数量
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // skuId
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

}
