package axindata

// FscSalesInfoApiDto 结构体
type FscSalesInfoApiDto struct {
	// 售卖范围 WEB:PC分销平台 WECHAT:阿信乐赚微信小程序 WECHAT_PLAY:阿信乐游微信小程序 FLIGGY:飞猪自营渠道
	SalesRange string `json:"sales_range,omitempty" xml:"sales_range,omitempty"`
	// 付款方式 1: 一次性付款 2: 分段式付款（预定金+尾款）
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
}
