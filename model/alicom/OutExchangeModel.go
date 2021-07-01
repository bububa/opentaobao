package alicom

// OutExchangeModel 结构体
type OutExchangeModel struct {
	// 单位为分
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 淘宝nick
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 外部单号
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 单位为分
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 固定值填15
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 商户合作id
	Partner string `json:"partner,omitempty" xml:"partner,omitempty"`
	// 来源
	From string `json:"from,omitempty" xml:"from,omitempty"`
	// 兑换的购物券面额（单位分）
	CouponFace string `json:"coupon_face,omitempty" xml:"coupon_face,omitempty"`
	// 扩展属性
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
}
