package lstpos

// CashierPayDetailDto 
type CashierPayDetailDto struct {
    // 备注信息
    DescInfo   string `json:"desc_info,omitempty" xml:"desc_info,omitempty"`
    // 买家支付/退款账号Id 若payType=alipay ,需要传递 买家的支付宝用户id；其它支付方式可不传
    PayAccount   string `json:"pay_account,omitempty" xml:"pay_account,omitempty"`
    // 创建时间 精确到毫秒
    GmtCreate   int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 支付状态    请求失败:requestFail    支付失败:payFail  支付超时:payTimeOut    支付成功:paySuccess  退款成功:refundSuccess 退款失败:refundFail
    PayStatus   string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    // 支付金额 单位：分 缺省值:0
    PayAmount   int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
    // 支付类型 支付宝:alipay  微信:wechat 现金:cashPay， 其它:other
    PayType   string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 支付流水号  正向支付:支付流水号; 逆向支付:退款流水号;
    PayFlowNum   string `json:"pay_flow_num,omitempty" xml:"pay_flow_num,omitempty"`
    // 业务类型  正向支付(支付)：0  缺省为：0   逆向支付(退款)：1
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
