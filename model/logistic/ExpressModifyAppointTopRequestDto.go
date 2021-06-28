package logistic

// ExpressModifyAppointTopRequestDto 
/* model for simplify = false
type ExpressModifyAppointTopRequestDto struct {

    // 应到达日期
    
    ScDate   string `json:"sc_date,omitempty"`
    

    // 子交易单号
    
    SubTradeIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sub_trade_ids,omitempty"`
    

    // 交易号
    
    TradeId   string `json:"trade_id,omitempty"`
    

    // 卖家Id
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 收货人电话
    
    ReceiverMobile   string `json:"receiver_mobile,omitempty"`
    

    // 改约日期
    
    OsDate   string `json:"os_date,omitempty"`
    

    // 收货人姓名
    
    ReceiverName   string `json:"receiver_name,omitempty"`
    

    // 扩展字段
    
    Feature   string `json:"feature,omitempty"`
    

    // 外部订单号
    
    OutOrderCode   string `json:"out_order_code,omitempty"`
    

    // 收货人地址
    
    ReceiverAddress   string `json:"receiver_address,omitempty"`
    

}
*/

// ExpressModifyAppointTopRequestDto 
type ExpressModifyAppointTopRequestDto struct {

    // 应到达日期
    ScDate   string `json:"sc_date,omitempty"`

    // 子交易单号
    SubTradeIds   []string `json:"sub_trade_ids,omitempty"`

    // 交易号
    TradeId   string `json:"trade_id,omitempty"`

    // 卖家Id
    SellerId   int64 `json:"seller_id,omitempty"`

    // 收货人电话
    ReceiverMobile   string `json:"receiver_mobile,omitempty"`

    // 改约日期
    OsDate   string `json:"os_date,omitempty"`

    // 收货人姓名
    ReceiverName   string `json:"receiver_name,omitempty"`

    // 扩展字段
    Feature   string `json:"feature,omitempty"`

    // 外部订单号
    OutOrderCode   string `json:"out_order_code,omitempty"`

    // 收货人地址
    ReceiverAddress   string `json:"receiver_address,omitempty"`

}
