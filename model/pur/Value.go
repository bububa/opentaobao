package pur

// Value 
type Value struct {

    // 订单编号
    
    PoNo   string `json:"po_no,omitempty" xml:"po_no,omitempty"`
    

    // 订单生效时间
    
    EffectTime   string `json:"effect_time,omitempty" xml:"effect_time,omitempty"`
    

    // 合同编号
    
    ContractNo   string `json:"contract_no,omitempty" xml:"contract_no,omitempty"`
    

    // 币种
    
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    

    // 订单金额
    
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    

    // 已发货金额
    
    DeliverAmount   string `json:"deliver_amount,omitempty" xml:"deliver_amount,omitempty"`
    

    // 已接收金额
    
    ReceiveAmount   string `json:"receive_amount,omitempty" xml:"receive_amount,omitempty"`
    

    // 已开票金额
    
    BillAmount   string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty"`
    

    // 已满足付款金额
    
    ToPayAmount   string `json:"to_pay_amount,omitempty" xml:"to_pay_amount,omitempty"`
    

    // 已付款金额
    
    PaymentAmount   string `json:"payment_amount,omitempty" xml:"payment_amount,omitempty"`
    

    // 状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 开始日期
    
    BeginDate   string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
    

    // 结束日期
    
    EndDate   string `json:"end_date,omitempty" xml:"end_date,omitempty"`
    

}
