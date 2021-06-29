package tvpay

// PartnerPayResultDo 
type PartnerPayResultDo struct {
    // 金额
    FundMoney   string `json:"fund_money,omitempty" xml:"fund_money,omitempty"`
    // 金额构成
    FundMoneyCode   string `json:"fund_money_code,omitempty" xml:"fund_money_code,omitempty"`
    // 手机号
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 订单号
    OrderNo   string `json:"order_no,omitempty" xml:"order_no,omitempty"`
    // 支付结果码
    PayCode   string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
    // 支付模式码
    PayMode   string `json:"pay_mode,omitempty" xml:"pay_mode,omitempty"`
}
