package mos

// OnsiteTradePayResponse 
type OnsiteTradePayResponse struct {

    // 喵街交易凭证号。必然返回
    
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    

    // 原支付请求的商户订单号。必然返回
    
    OutTradeNo   string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
    

    // 消费者在喵街的昵称，将用*号屏蔽部分内容。必然返回
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 订单总金额，单位为分。必然返回
    
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    

    // 交易状态。必然返回。取值：WAIT_FOR_CONFIRM，待确认，此时不能申请退款，可以撤销；WAIT_BUYER_PAY，等待消费者付款	，此时不能申请退款，可以撤销；TRADE_SUCCESS，付款成功，此时可以申请退款，可以撤销（自动申请退款）；TRADE_FINISHED，交易成功，此时可以申请退款，可以撤销（自动申请退款）；TRADE_CLOSED，交易关闭，此时不能申请退款，不能撤销
    
    TradeStatus   string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
    

    // 码来源，取值：MJ、M_TAO、ALIPAY
    
    AuthCodeSource   string `json:"auth_code_source,omitempty" xml:"auth_code_source,omitempty"`
    

}
