package alsc

// PayDetailInfo 
type PayDetailInfo struct {
    // 业务操作人id
    OutOperatorId   string `json:"out_operator_id,omitempty" xml:"out_operator_id,omitempty"`
    // 业务操作人名称
    OutOperatorName   string `json:"out_operator_name,omitempty" xml:"out_operator_name,omitempty"`
    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    // 支付方式id
    OutPayChannelId   string `json:"out_pay_channel_id,omitempty" xml:"out_pay_channel_id,omitempty"`
    // 支付方式名称
    OutPayChannelName   string `json:"out_pay_channel_name,omitempty" xml:"out_pay_channel_name,omitempty"`
    // 支付方式code
    OutPayCode   string `json:"out_pay_code,omitempty" xml:"out_pay_code,omitempty"`
    // 外部支付明细ID
    OutPayDetailId   string `json:"out_pay_detail_id,omitempty" xml:"out_pay_detail_id,omitempty"`
    // 外部交易单号，支付通道的支付单号
    OutTradeNo   string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
    // 支付账号，如微信open_id，支付宝buyer_id
    PayAccountId   string `json:"pay_account_id,omitempty" xml:"pay_account_id,omitempty"`
    // 支付卡号
    PayCardId   string `json:"pay_card_id,omitempty" xml:"pay_card_id,omitempty"`
    // 支付金额，单位“分”
    PayFee   int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
    // 支付来源 HANDHELD_DEVICES手持设备 SAMPLE_ORDER扫码点餐 POS POS端 MSTORE 移动门店 THIRDPARTY第三方 OTHER 其他
    PaySource   string `json:"pay_source,omitempty" xml:"pay_source,omitempty"`
    // 支付状态 SUCCESS 成功  FAIL 失败
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 支付完成时间
    SuccessTime   int64 `json:"success_time,omitempty" xml:"success_time,omitempty"`
}
