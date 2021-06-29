package iotticket

// AssignMaintenancePersonnelTopRequest 
type AssignMaintenancePersonnelTopRequest struct {

    // 操作人联系方式
    
    OperatorPhone   string `json:"operator_phone,omitempty" xml:"operator_phone,omitempty"`
    

    // 操作人编码
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 操作人名称
    
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    

    // 服务商唯一编码
    
    SpCode   string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
    

    // 工单Id
    
    TicketId   int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
    

    // 维修项
    
    IotMaintainPlanItemList   []IotMaintainPlanItemTopRequest `json:"iot_maintain_plan_item_list,omitempty" xml:"iot_maintain_plan_item_list,omitempty"`
    

    // 运维方案：SEND_BACK_AND_SEND_OUT-客户寄回服务商寄出；SEND_OUT-服务商寄出；ONSITE-上门服务
    
    MaintainAbilities   string `json:"maintain_abilities,omitempty" xml:"maintain_abilities,omitempty"`
    

    // 客户寄回设备 服务商收件人名称
    
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    

    // 其它费用
    
    OtherFee   string `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
    

    // 客户寄回设备 服务商收货地址
    
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    

    // 扩展字段
    
    Features   string `json:"features,omitempty" xml:"features,omitempty"`
    

    // 维修方式（需要映射）
    
    MaintainMethod   string `json:"maintain_method,omitempty" xml:"maintain_method,omitempty"`
    

    // 客户寄回设备 服务商联系方式
    
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    

    // 支付方式：payBefore-维修前付费;payAfter-维修后付费;noNeedPay-无需付费
    
    PayMethod   string `json:"pay_method,omitempty" xml:"pay_method,omitempty"`
    

    // 费用描述
    
    FeeRemark   string `json:"fee_remark,omitempty" xml:"fee_remark,omitempty"`
    

    // 保内保外（需要映射）
    
    WarrantyType   string `json:"warranty_type,omitempty" xml:"warranty_type,omitempty"`
    

    // 事件类型（需要映射）
    
    EventTypeList   []string `json:"event_type_list,omitempty" xml:"event_type_list>string,omitempty"`
    

}
