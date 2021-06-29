package train

// AgentAgreeChangeParam 
type AgentAgreeChangeParam struct {
    // 申请单id
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 支付宝交易流水号
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
    // 扩展参数
    ExtendParam   string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 代理商id
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    // 主订单
    MainBizOrderId   int64 `json:"main_biz_order_id,omitempty" xml:"main_biz_order_id,omitempty"`
    // 代理商改签票信息
    Tickets   []ChangeTicketInfo `json:"tickets,omitempty" xml:"tickets>change_ticket_info,omitempty"`
    // 支付宝账号
    AlipayAccount   string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
}
