package flight

// ModifyDetailDto 
type ModifyDetailDto struct {

    // 申请单号
    
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    

    // 申请原因
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

    // 国际国内标识
    
    DomesticIntl   int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
    

    // 关联飞猪订单号
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 改签数据
    
    ChangeList   []ChangeList `json:"change_list,omitempty" xml:"change_list,omitempty"`
    

    // sla
    
    Sla   string `json:"sla,omitempty" xml:"sla,omitempty"`
    

    // 币种
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 申请时间
    
    ApplyTime   string `json:"apply_time,omitempty" xml:"apply_time,omitempty"`
    

    // 改签单状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 佣金
    
    Commission   int64 `json:"commission,omitempty" xml:"commission,omitempty"`
    

    // 是否自愿
    
    Voluntary   int64 `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
    

    // 拒绝原因
    
    RefuseReason   string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
    

    // 订单标签
    
    Tags   []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
    

}
