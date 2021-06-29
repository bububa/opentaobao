package westcrm

// OrderVo 
type OrderVo struct {
    // 账单金额
    OrderPay   string `json:"order_pay,omitempty" xml:"order_pay,omitempty"`
    // 账单id
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 消费时间
    OrderTime   string `json:"order_time,omitempty" xml:"order_time,omitempty"`
}
