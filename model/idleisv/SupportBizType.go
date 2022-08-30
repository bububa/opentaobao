package idleisv

// SupportBizType 结构体
type SupportBizType struct {
	// 0 已验货-不入仓，1 已验货-入仓，3 寄卖，4 品牌免检，5  官方自营，7 品牌直营，8 专业认证，9  信誉担保
	ItemBizType int64 `json:"item_biz_type,omitempty" xml:"item_biz_type,omitempty"`
}
