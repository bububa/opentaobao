package jst

// ChildOcOrder 
type ChildOcOrder struct {

    // 分帐金额
    AllocatePaymentAmount   int64 `json:"allocate_payment_amount,omitempty"`

    // 订单编号
    SourceTradeNo   string `json:"source_trade_no,omitempty"`

    // 主扣款账号
    MainAllocationPaymentAcount   int64 `json:"main_allocation_payment_acount,omitempty"`

    // 订单来源渠道TB,或者TMALL
    SourceTradeChannel   string `json:"source_trade_channel,omitempty"`

    // 分账ruleId,调用创建规则获取
    AllocatePaymentRuleId   int64 `json:"allocate_payment_rule_id,omitempty"`

    // 分账参与方
    AllocatePaymentParticipants   string `json:"allocate_payment_participants,omitempty"`

    // 子订单id
    Id   int64 `json:"id,omitempty"`

    // 父订单id
    ParentId   int64 `json:"parent_id,omitempty"`

}
