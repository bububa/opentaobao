package trade

// TradeConfirmFee 
type TradeConfirmFee struct {

    // 确认收货的金额(不包含邮费)。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    ConfirmFee   string `json:"confirm_fee,omitempty" xml:"confirm_fee,omitempty"`
    

    // 需确认收货的邮费，当不是最后一笔收货或者没有邮费时是0.00。精确到2位小数;单位:元。如:200.07，表示:200元7分
    
    ConfirmPostFee   string `json:"confirm_post_fee,omitempty" xml:"confirm_post_fee,omitempty"`
    

    // 是否是最后一笔订单（只对子订单有效，当其他子订单都是交易完成时，返回true，否则false）
    
    IsLastOrder   bool `json:"is_last_order,omitempty" xml:"is_last_order,omitempty"`
    

}
