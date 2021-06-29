package idleisv

// IdleItemApiSkuDo 
type IdleItemApiSkuDo struct {

    // skuId
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 价格，单位分
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 库存
    
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    

    // 销售属性列表(最多2个销售属性,每一个的属性值个数为2～10)
    
    PropertyList   []IdleItemApiPvPairDo `json:"property_list,omitempty" xml:"property_list,omitempty"`
    

}
