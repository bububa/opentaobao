package iotticket

// MaintenanceInfo 
type MaintenanceInfo struct {

    // 维修方式
    
    MaintainMethod   string `json:"maintain_method,omitempty" xml:"maintain_method,omitempty"`
    

    // 修理方式
    
    MaintainAbilities   string `json:"maintain_abilities,omitempty" xml:"maintain_abilities,omitempty"`
    

    // 支付方式：payBefore-维修前付费;payAfter-维修后付费;noNeedPay-无需付费
    
    PayMethod   string `json:"pay_method,omitempty" xml:"pay_method,omitempty"`
    

    // 其它费用
    
    OtherFee   string `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
    

    // 费用描述
    
    FeeRemark   string `json:"fee_remark,omitempty" xml:"fee_remark,omitempty"`
    

    // 维修配件信息
    
    PartItemList   []PartItemList `json:"part_item_list,omitempty" xml:"part_item_list,omitempty"`
    

    // 承保类型
    
    WarrantyType   string `json:"warranty_type,omitempty" xml:"warranty_type,omitempty"`
    

    // 事件类型列表
    
    EventTypeList   []string `json:"event_type_list,omitempty" xml:"event_type_list>string,omitempty"`
    

}
