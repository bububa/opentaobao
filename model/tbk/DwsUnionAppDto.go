package tbk

// DwsUnionAppDto 结构体
type DwsUnionAppDto struct {
	// 付款预估收入。指买家付款金额为基数，预估您可能获得的收入。因买家退款等原因，可能与结算预估收入不一致
	Prepubsharefee string `json:"pre_pub_share_fee,omitempty" xml:"pre_pub_share_fee,omitempty"`
	// 支付金额
	Alipayamt string `json:"alipay_amt,omitempty" xml:"alipay_amt,omitempty"`
	// 人群内成交人数
	Alipayusernuminucrowd int64 `json:"alipay_user_num_in_ucrowd,omitempty" xml:"alipay_user_num_in_ucrowd,omitempty"`
	// 支付人数
	Alipayusernum int64 `json:"alipay_user_num,omitempty" xml:"alipay_user_num,omitempty"`
	// 支付订单数
	Alipaynum int64 `json:"alipay_num,omitempty" xml:"alipay_num,omitempty"`
	// 点击人数
	Clickuvunion int64 `json:"click_uv_union,omitempty" xml:"click_uv_union,omitempty"`
	// 点击数
	Clickpvunion int64 `json:"click_pv_union,omitempty" xml:"click_pv_union,omitempty"`
}
