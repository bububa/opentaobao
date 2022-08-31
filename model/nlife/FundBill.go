package nlife

// FundBill 结构体
type FundBill struct {
	// 资金渠道的id:     * 1. type=CASH：为空；      * 2. type=UNIONPAY：银联流水号；      * 3. type=ALIPAY：支付宝订单号；      * 4. type=WECHAT_PAY：微信支付订单号；      * 6. type=CUSTOM_PROMOTION：零售商自有优惠id。      * 7. type=CUSTOM_PREPAY_CARD：零售商自有储值卡id      * 8. type=MALING：为空
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 资金渠道类型:     CASH(&#34;现金支付&#34;),     UNIONPAY(&#34;银联刷卡支付&#34;),     ALIPAY(&#34;支付宝支付&#34;),     WECHAT_PAY(&#34;微信支付&#34;),     PROMOTION(&#34;零售+平台的优惠&#34;),     CUSTOM_PROMOTION(&#34;零售商自有优惠&#34;),     CUSTOM_PREPAY_CARD(&#34;零售商自有储值卡&#34;),     MALING(&#34;支付时抹零&#34;);
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 用户ID
	BuyerId string `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 买家类型： TAOBAO_OPENID(&#34;经阿里百川淘宝登陆获取到的用户ID&#34;),     ALIPAY_OPENID(&#34;经蚂蚁金服开放平台支付宝App登陆获得的用户ID&#34;),     WECHAT_OPENID(&#34;经微信开放平台微信App登陆获得的用户id&#34;),     PHONE_NUMBER(&#34;通过手机号码登陆&#34;),     APP_USERID(&#34;商户自由的用户ID&#34;),     ANONYMOUS_USER(&#34;匿名用户&#34;);
	BuyerIdType string `json:"buyer_id_type,omitempty" xml:"buyer_id_type,omitempty"`
	// 标题：     * 1. type=CASH：现金支付；      * 2. type=UNIONPAY：银联刷卡支付；      * 3. type=ALIPAY：支付宝支付；      * 4. type=WECHAT_PAY：微信支付；      * 5. type=PROMOTION: 零售+平台的优惠名称；      * 6. type=CUSTOM_PROMOTION：零售商自有优惠名称；      * 7. type=CUSTOM_PREPAY_CARD：零售商自有储值卡；      * 8. type=MALING：抹零      *
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 货币种类：    CYN(&#34;人民币&#34;),     USD(&#34;美元&#34;),     HKD(&#34;港币&#34;);
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 扩展参数，JSON格式
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 附加数据，在查询订单中原样返回，该字段主要用于商户携带订单的自定义数据
	Attachment string `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// 金额/优惠抵扣金额，单位：分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}
