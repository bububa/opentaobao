package flight

// RefundDetailDto 
/* model for simplify = false
type RefundDetailDto struct {

    // 申请单号
    
    ApplyId   string `json:"apply_id,omitempty"`
    

    // 申请原因类型
    
    ApplyReasonType   int64 `json:"apply_reason_type,omitempty"`
    

    // 申请原因
    
    Reason   string `json:"reason,omitempty"`
    

    // 店铺id
    
    AgentId   int64 `json:"agent_id,omitempty"`
    

    // 国际国内标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`
    

    // 关联飞猪订单号
    
    OrderId   string `json:"order_id,omitempty"`
    

    // 退票数据集
    
    RefundList  struct {
        RefundList  []RefundList `json:"refund_list,omitempty"`
    } `json:"refund_list,omitempty"`
    

    // sla
    
    Sla   string `json:"sla,omitempty"`
    

    // 交易币种
    
    Currency   string `json:"currency,omitempty"`
    

    // 申请时间
    
    ApplyTime   string `json:"apply_time,omitempty"`
    

    // 退票状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 自愿标识
    
    Voluntary   int64 `json:"voluntary,omitempty"`
    

    // 拒绝原因
    
    RefuseReason   string `json:"refuse_reason,omitempty"`
    

    // 补退单标识(1是补退单)
    
    Supplement   int64 `json:"supplement,omitempty"`
    

    // 紧急标识
    
    Tags  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tags,omitempty"`
    

}
*/

// RefundDetailDto 
type RefundDetailDto struct {

    // 申请单号
    ApplyId   string `json:"apply_id,omitempty"`

    // 申请原因类型
    ApplyReasonType   int64 `json:"apply_reason_type,omitempty"`

    // 申请原因
    Reason   string `json:"reason,omitempty"`

    // 店铺id
    AgentId   int64 `json:"agent_id,omitempty"`

    // 国际国内标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty"`

    // 关联飞猪订单号
    OrderId   string `json:"order_id,omitempty"`

    // 退票数据集
    RefundList   []RefundList `json:"refund_list,omitempty"`

    // sla
    Sla   string `json:"sla,omitempty"`

    // 交易币种
    Currency   string `json:"currency,omitempty"`

    // 申请时间
    ApplyTime   string `json:"apply_time,omitempty"`

    // 退票状态
    Status   int64 `json:"status,omitempty"`

    // 自愿标识
    Voluntary   int64 `json:"voluntary,omitempty"`

    // 拒绝原因
    RefuseReason   string `json:"refuse_reason,omitempty"`

    // 补退单标识(1是补退单)
    Supplement   int64 `json:"supplement,omitempty"`

    // 紧急标识
    Tags   []string `json:"tags,omitempty"`

}
