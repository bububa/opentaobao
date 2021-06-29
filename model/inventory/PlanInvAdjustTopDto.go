package inventory

// PlanInvAdjustTopDto 
type PlanInvAdjustTopDto struct {

    // 品id，前端宝贝id，或者后端货品id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 品的类型，1-前端宝贝，2-后端货品
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // 计划库存增量编辑的详细信息
    
    AdjustDetailList   []PlanInvAdjustTopDetailDto `json:"adjust_detail_list,omitempty" xml:"adjust_detail_list,omitempty"`
    

    // 设置计划的品的skuid。对于货品，是0.
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

}
