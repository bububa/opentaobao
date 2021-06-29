package alitripmerchant

// ValidateOrderVo 
type ValidateOrderVo struct {
    // 实际需要支付金额，含税
    TotalAmount   string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 总税价
    TotalTax   string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
    // 价格变化描述
    AmountChangedDisplay   string `json:"amount_changed_display,omitempty" xml:"amount_changed_display,omitempty"`
    // 价格发生变化
    IsAmountChanged   bool `json:"is_amount_changed,omitempty" xml:"is_amount_changed,omitempty"`
    // 每日价格
    DailyPriceList   []DailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>daily_price,omitempty"`
    // 订单号
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
}
