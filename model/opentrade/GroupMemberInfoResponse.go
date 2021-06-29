package opentrade

// GroupMemberInfoResponse 
type GroupMemberInfoResponse struct {

    // 1-未付款，2-已付款，4-已退款，6-交易成功，8-交易关闭，null表示用户未参团
    
    PayStatus   int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    

    // 用户openId
    
    OpenUserId   string `json:"open_user_id,omitempty" xml:"open_user_id,omitempty"`
    

    // 用户参团订单号
    
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

}
