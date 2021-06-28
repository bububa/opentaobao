package refund

// SyncIdentifyRefundCaseDto 
/* model for simplify = false
type SyncIdentifyRefundCaseDto struct {

    // 子订单ID
    
    DetailOrderId   int64 `json:"detail_order_id,omitempty"`
    

    // 数据发生时间绝对秒数，如鉴定工单创建时间、鉴定工单完结时间
    
    OccurTime   int64 `json:"occur_time,omitempty"`
    

    // 工单操作类型，1：开启，2：完结
    
    OperateType   int64 `json:"operate_type,omitempty"`
    

    // 鉴定工单操作备注
    
    OperateTips   string `json:"operate_tips,omitempty"`
    

    // 鉴定工单ID
    
    OuterCaseId   string `json:"outer_case_id,omitempty"`
    

    // 退款ID
    
    RefundId   int64 `json:"refund_id,omitempty"`
    

    // 扩展属性，json格式
    
    ExtAttrs   string `json:"ext_attrs,omitempty"`
    

}
*/

// SyncIdentifyRefundCaseDto 
type SyncIdentifyRefundCaseDto struct {

    // 子订单ID
    DetailOrderId   int64 `json:"detail_order_id,omitempty"`

    // 数据发生时间绝对秒数，如鉴定工单创建时间、鉴定工单完结时间
    OccurTime   int64 `json:"occur_time,omitempty"`

    // 工单操作类型，1：开启，2：完结
    OperateType   int64 `json:"operate_type,omitempty"`

    // 鉴定工单操作备注
    OperateTips   string `json:"operate_tips,omitempty"`

    // 鉴定工单ID
    OuterCaseId   string `json:"outer_case_id,omitempty"`

    // 退款ID
    RefundId   int64 `json:"refund_id,omitempty"`

    // 扩展属性，json格式
    ExtAttrs   string `json:"ext_attrs,omitempty"`

}
