package inventory

// InvUnifyPlanTopQuery 
type InvUnifyPlanTopQuery struct {

    // 商品或者货品的id，计划建在哪，就用哪个id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // skuid。如果是货品，则skuid是0
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // item_id的类型，1是前端宝贝，2是后端货品
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // 生成计划库存的外部单据号
    
    PlanOrderId   string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
    

}
