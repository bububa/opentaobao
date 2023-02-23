package alitripmerchant

// DerbyVoucherCardPayDto 结构体
type DerbyVoucherCardPayDto struct {
	// 实付金额
	PaidInAmount string `json:"paid_in_amount,omitempty" xml:"paid_in_amount,omitempty"`
	// 商品权益编码
	VoucherCardCode string `json:"voucher_card_code,omitempty" xml:"voucher_card_code,omitempty"`
	// 商品类型APLUSP:Premium卡。APLUSB:Basic卡。
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 是否接受会员条款(必选)
	AcceptedTandC string `json:"accepted_tand_c,omitempty" xml:"accepted_tand_c,omitempty"`
	// 是否勾选邮箱条款
	OptinAll string `json:"optin_all,omitempty" xml:"optin_all,omitempty"`
	// 订单id(传了就是已存在订单，重复唤起支付)
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 是否勾选短信条款
	OptinSMSALL bool `json:"optin_s_m_s_a_l_l,omitempty" xml:"optin_s_m_s_a_l_l,omitempty"`
}
