package mos

// Product 
type Product struct {

    // 分摊金额
    Amount   string `json:"amount,omitempty"`

    // 支付方式行号
    Fkfsorder   string `json:"fkfsorder,omitempty"`

    // 订单号或券号（支付宝订单号）
    Orderid   string `json:"orderid,omitempty"`

    // 商品行号
    Comorder   string `json:"comorder,omitempty"`

    // 大支付方式
    Payment   string `json:"payment,omitempty"`

    // 小支付方式
    Subpayment   string `json:"subpayment,omitempty"`

}
