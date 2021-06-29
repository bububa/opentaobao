package nlife

// ItemQueryDOList 
type ItemQueryDOList struct {

    // 商品ID(与outer_id不能同时为空)
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 商品在外部商家的编码(与Item_id不能同时为空)
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

    // skuId列表-可查询指定的sku
    
    SkuIdList   []int64 `json:"sku_id_list,omitempty" xml:"sku_id_list>int64,omitempty"`
    

}
