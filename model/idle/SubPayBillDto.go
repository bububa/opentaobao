package idle

// SubPayBillDTO 
type SubPayBillDTO struct {
    // 金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 代扣错误码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 支付状态描述
    PayStatusDesc   string `json:"pay_status_desc,omitempty" xml:"pay_status_desc,omitempty"`
    // 代扣错误信息
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // 计划id
    PlanId   string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
    // 支付状态
    PayStatus   string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    // 支付宝流水号
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
}
