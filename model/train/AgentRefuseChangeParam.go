package train

// AgentRefuseChangeParam 
type AgentRefuseChangeParam struct {

    // 改签申请单id
    ApplyId   int64 `json:"apply_id,omitempty"`

    // 扩展参数
    ExtendParam   string `json:"extend_param,omitempty"`

    // 拒绝原因类型，0 其他、1 余票不足、2 已取票、3 已线下退票
    RefuseType   int64 `json:"refuse_type,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 代理商id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 订单id
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty"`

}
