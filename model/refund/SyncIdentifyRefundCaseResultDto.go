package refund

// SyncIdentifyRefundCaseResultDto 
/* model for simplify = false
type SyncIdentifyRefundCaseResultDto struct {

    // 子订单ID
    
    DetailOrderId   int64 `json:"detail_order_id,omitempty"`
    

    // 数据发生时间绝对秒数，如写入鉴定结果的时间
    
    OccurTime   int64 `json:"occur_time,omitempty"`
    

    // 鉴定工单ID
    
    OuterCaseId   string `json:"outer_case_id,omitempty"`
    

    // 结果类型，1：可退，2：可换，3：不通过
    
    ResultType   int64 `json:"result_type,omitempty"`
    

    // 退款ID
    
    RefundId   int64 `json:"refund_id,omitempty"`
    

    // 鉴定结果提示
    
    ResultTips   string `json:"result_tips,omitempty"`
    

    // 扩展属性，json格式
    
    ExtAttrs   string `json:"ext_attrs,omitempty"`
    

}
*/

// SyncIdentifyRefundCaseResultDto 
type SyncIdentifyRefundCaseResultDto struct {

    // 子订单ID
    DetailOrderId   int64 `json:"detail_order_id,omitempty"`

    // 数据发生时间绝对秒数，如写入鉴定结果的时间
    OccurTime   int64 `json:"occur_time,omitempty"`

    // 鉴定工单ID
    OuterCaseId   string `json:"outer_case_id,omitempty"`

    // 结果类型，1：可退，2：可换，3：不通过
    ResultType   int64 `json:"result_type,omitempty"`

    // 退款ID
    RefundId   int64 `json:"refund_id,omitempty"`

    // 鉴定结果提示
    ResultTips   string `json:"result_tips,omitempty"`

    // 扩展属性，json格式
    ExtAttrs   string `json:"ext_attrs,omitempty"`

}
