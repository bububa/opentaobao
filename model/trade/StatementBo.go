package trade

// StatementBo 结构体
type StatementBo struct {
	// 对账单号
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 渠道编码 WX_APPLET(&#34;10001&#34;, &#34;微信小程序&#34;),     WX_OFFICIAL_SERVER(&#34;10002&#34;, &#34;微信服务号/微信商城&#34;),     ALIPAY_APPLET(&#34;20001&#34;, &#34;支付宝小程序&#34;),     OFFLINE_MEMBER_CARD(&#34;30001&#34;, &#34;线下会员卡&#34;),     MERCHANT_APP(&#34;40001&#34;, &#34;商家自有app&#34;),     TXD_APP(&#34;50001&#34;, &#34;淘鲜达app&#34;),     TXD_SELF_POS(&#34;60001&#34;, &#34;淘鲜达自助pos&#34;),     TXD_ARTIFACT_POS(&#34;60001&#34;, &#34;淘鲜达人工pos&#34;);
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 订单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 外部会员Id
	OpenMemberId string `json:"open_member_id,omitempty" xml:"open_member_id,omitempty"`
	// 渠道订单Id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单类型     NORMAL(0, &#34;正向&#34;),     REFUND(1, &#34;逆向&#34;),     REVOKED(2, &#34;撤销&#34;);
	OrderType int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 支付方式 ALIPAY(1, &#34;支付宝&#34;),     WECHAT(2, &#34;微信&#34;),     CASH(3, &#34;现金&#34;),     VALUE_CARD(4, &#34;储值卡&#34;),     OTHER(99, &#34;其他&#34;)
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
}
