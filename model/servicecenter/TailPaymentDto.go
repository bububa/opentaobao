package servicecenter

// TailPaymentDTO 
type TailPaymentDTO struct {
    // 买家id，不需要传入
    BuyerId   int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
    // 月供，单位分，如果是购买车辆，分期付尾款，则必须
    MonthlyPay   int64 `json:"monthly_pay,omitempty" xml:"monthly_pay,omitempty"`
    // 分期月份，如果是购买车辆，分期付尾款，则必须
    Months   int64 `json:"months,omitempty" xml:"months,omitempty"`
    // 处置名字，支持3种：归还车辆; 购买车辆，一次性支付尾款; 购买车辆，分期付尾款
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 尾款金额，单位分，如果是购买车辆，一次性支付尾款必填
    TailAmount   int64 `json:"tail_amount,omitempty" xml:"tail_amount,omitempty"`
}
