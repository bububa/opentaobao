package flight

// RefundDetailDto 
type RefundDetailDto struct {
    // 申请单号
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 申请原因类型
    ApplyReasonType   int64 `json:"apply_reason_type,omitempty" xml:"apply_reason_type,omitempty"`
    // 申请原因
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    // 店铺id
    AgentId   int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
    // 国际国内标识
    DomesticIntl   int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
    // 关联飞猪订单号
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 退票数据集
    RefundList   []RefundList `json:"refund_list,omitempty" xml:"refund_list>refund_list,omitempty"`
    // sla
    Sla   string `json:"sla,omitempty" xml:"sla,omitempty"`
    // 交易币种
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 申请时间
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    // 退票状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 自愿标识
    Voluntary   int64 `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
    // 拒绝原因
    RefuseReason   string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
    // 补退单标识(1是补退单)
    Supplement   int64 `json:"supplement,omitempty" xml:"supplement,omitempty"`
    // 紧急标识
    Tags   []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
}
