package promotion

// CouponTemplateCommonConfig 结构体
type CouponTemplateCommonConfig struct {
	// 申请渠道 anonymousOffline
	ApplyChannels []string `json:"apply_channels,omitempty" xml:"apply_channels>string,omitempty"`
	// 优惠券活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 优惠券名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 优惠券发放方式 ANONYMOUS(&#34;anonymous&#34;,&#34;匿名券&#34;), REGISTERED(&#34;registered&#34;,&#34;记名券&#34;),
	SendType string `json:"send_type,omitempty" xml:"send_type,omitempty"`
	// 模板状态 NORMAL(1, &#34;有效&#34;), DELETE(-1, &#34;删除&#34;), ENDING(0, &#34;结束领取&#34;), NOUSE(-2, &#34;无效&#34;), WDK_COUPON_DRAFT(-999, &#34;草稿&#34;),
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 优惠券类型 UNIT_PRICE(10,&#34;unitPrice&#34;,&#34;单品定价券&#34;), FULL_AMOUNT_REDUCE(11, &#34;fullAmountReduce&#34;, &#34;满元减券&#34;), FULL_AMOUNT_DISCOUNT(12, &#34;fullAmountDiscount&#34;, &#34;满元折券&#34;), FULL_COUNT_REDUCE(13, &#34;fullCountReduce&#34;, &#34;满件减券&#34;), FULL_COUNT_DISCOUNT(14, &#34;fullCountDiscount&#34;, &#34;满件折券&#34;), VOUCHER(15, &#34;voucher&#34;, &#34;抵用券&#34;),
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
