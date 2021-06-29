package inventory

// PlanInstanceInvResultDto 
type PlanInstanceInvResultDto struct {

    // 品id
    
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 品类型，1是前端宝贝，2是后端货品
    
    ItemType   int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
    

    // skuid
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 当前计划的实时剩余可售数量
    
    SellableQuantity   int64 `json:"sellable_quantity,omitempty" xml:"sellable_quantity,omitempty"`
    

    // 当前计划的实时预扣数量
    
    WithholdingQuantity   int64 `json:"withholding_quantity,omitempty" xml:"withholding_quantity,omitempty"`
    

    // 当前计划的实时占用数量
    
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
    

    // 计划的开始销售时间
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 计划的结束销售时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

    // 策略
    
    Strategy   *StrategyRuleTopDto `json:"strategy,omitempty" xml:"strategy,omitempty"`
    

    // 生成计划库存的外部单据号
    
    PlanOrderId   string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
    

    // 计划id，平台内为planOrderId生成的计划id
    
    PlanInstanceId   int64 `json:"plan_instance_id,omitempty" xml:"plan_instance_id,omitempty"`
    

    // 最后一次全量设置的计划库存值。如果是增量编辑库存，settingQuantity不会变。这个值仅做参考，主要看前面的实时剩余可售数量sellable_quantity
    
    SettingQuantity   int64 `json:"setting_quantity,omitempty" xml:"setting_quantity,omitempty"`
    

    // 计划的履约仓信息
    
    PromiseList   []PlanPromiseTopDto `json:"promise_list,omitempty" xml:"promise_list,omitempty"`
    

    // 计划生效的前端宝贝列表，如果是货品关联的所有前端都生效，则这个对象是空
    
    RelationList   []PlanItemRelationTopDto `json:"relation_list,omitempty" xml:"relation_list,omitempty"`
    

}
